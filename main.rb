# The file in Lambda which converts audio to text

require 'json'
require 'aws-sdk-transcribeservice'
require 'aws-sdk-cloudwatchlogs'
require 'aws-sdk-s3'

def lambda_handler(event:, context:)
    # Get the bucket name and object key from the event data
  object_key = event['Records'][0]['s3']['object']['key']

  # Create an instance of the S3 client
  s3_client = Aws::S3::Client.new(region: 'eu-north-1')

  # Get the object from the S3 bucket
  response = s3_client.get_object(bucket: 'audiofilebucketprojekt', key: object_key)
  object_uri = "s3://audiofilebucketprojekt/#{object_key}"
  identifier = object_uri.match(/(\d{10})\./)[1]
  puts identifier
  
  media = {
    media_file_uri: object_uri
  }

  # Create an instance of the Transcribe client
  transcribe_client = Aws::TranscribeService::Client.new(region: 'eu-north-1')
  
  # Start the transcription job
  transcribe_client.start_transcription_job(
    transcription_job_name: "transcribe-job-#{identifier}",
    language_code: 'da-DK',
    media: media,
    output_bucket_name: 'audiofilebucketprojekt'
  )
  
  # Access the file data
  file_data = response.body.read
  #puts file_data
  puts 'Successfully accessed the file'

  # Return a response or perform any other actions
  
end
