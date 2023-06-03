// Please dont abuse these keys
AWS.config.update({
    accessKeyId: 'AKIA2PDQIEQEMGSVJY6P',
    secretAccessKey: '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'
});

const s3 = new AWS.S3({
    region: 'eu-north-1'
});

const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit)

function handleSubmit(event) {
    document.querySelector('#confirmButton').disabled = true;
    document.querySelector('.loader').style.display = 'flex';
    const fileInput = document.getElementById('audioFile');
    const file = fileInput.files[0];
    uploadFile(file);

    //getTranscription(7492254678)
    event.preventDefault();
}

function uploadFile(file) {
    const filename = file.name
    const uploadCode = generateRandomNumber();
    const lastDotIndex = filename.lastIndexOf("."); // Find the last occurrence of "."
    const newFilename = filename.substring(0, lastDotIndex) + uploadCode + filename.substring(lastDotIndex); // Insert code before the last "."

    const params = {
        Bucket: 'audiofilebucketprojekt',
        Key: `audiofiles/${newFilename}`,
        Body: file,
        ContentType: 'audio/wav'
    };

    s3.putObject(params, function (err, data) {
        if (err) {
            console.log(err, err.stack); // Handle error
        } else {
            console.log('File uploaded successfully');
            // file uploaded now try to get transcription
            getTranscription(uploadCode);
        }
    });
}

function getTranscription(uploadCode) {
    const params = {
        Bucket: 'audiofilebucketprojekt',
        Key: `transcribe-job-${uploadCode}.json`
    };
    s3.headObject(params, function (err, data) {
        if (err) {
            console.log('File does not exist.');
            setTimeout(() => { getTranscription(uploadCode) }, 2000)
        } else {
            // file exists
            console.log('File exists.');
            s3.getObject(params, function (err, data) {
                if (err) {
                    console.log(err, err.stack); // Handle error
                    // there was an error try again
                } else {
                    // Access the downloaded file data
                    const jsonFileData = JSON.parse(data.Body.toString());
                    const result = jsonFileData.results.transcripts[0].transcript
                    const transciptionText = document.querySelector('#textContent');
                    transciptionText.textContent = result;

                    document.querySelector('.loader').style.display = 'none';
                    document.querySelector('.copy-field').style.display = 'flex';
                }
            });
        }
    });
}

// genereates a random 10 digit number
function generateRandomNumber() {
    const min = 1000000000;
    const max = 9999999999;

    // Generate a random number within the specified range
    const randomNumber = Math.floor(Math.random() * (max - min + 1)) + min;

    return randomNumber;
}



const textField = document.getElementById('myTextField');
textField.addEventListener("click", copyText);
function copyText() {
    const textField = document.getElementById("textContent");
    //textField.select();
    //textField.setSelectionRange(0, 99999); /* For mobile devices */
    // Copy the text inside the text field
    navigator.clipboard.writeText(textField.innerText);
    alert("Text copied to clipboard!");
}

if(document.querySelector('#textContent').innerText !== '') {
    document.querySelector('.copy-field').style.display = 'flex'
}
const fileInput = document.getElementById('audioFile');
displayFileName();
fileInput.addEventListener("change", displayFileName);
function displayFileName() {
    const fileName = document.getElementById('fileName');
    console.log(fileInput.files.length)
    if (fileInput.files.length <= 0) {
        document.querySelector('#confirmButton').style.backgroundColor = 'rgb(114, 120, 135)';
        document.querySelector('#confirmButton').style.color = 'rgb(105, 114, 137)';
        document.querySelector('#confirmButton').disabled = true;
        document.querySelector('#confirmButton').style.boxShadow = 'none';
        document.querySelector('#confirmButton').style.cursor = 'default';
    } else {
        fileName.textContent = fileInput.files[0]?.name;
        document.querySelector('#confirmButton').style.backgroundColor = '';
        document.querySelector('#confirmButton').style.color = '';
        document.querySelector('#confirmButton').disabled = '';
        document.querySelector('#confirmButton').style.boxShadow = '';
        document.querySelector('#confirmButton').style.cursor = '';

    }
}
