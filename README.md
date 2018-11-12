# Copy to another S3 Bucket

AWS Lambda to copy an object created in one S3 bucket to another.

## How to configure?

Generate a binary file (linux OS):
```bash
GOOS=linux go build copy_to_another_s3_bucket.go
```

Generate a zip file:
```bash
zip copy_to_another_s3_bucket.zip ./copy_to_another_s3_bucket
```

AWS Lambda setup:

 1. Create an AWS Lambda function
 2. Set Runtime to `Go 1.x`
 3. Set Handler to `copy_to_another_s3_bucket`
 4. Upload the zip file in **Function package**
 5. Set to this function a **Role** with **AdministratorAccess**
 6. Add a S3 trigger (Event type: ObjectCreated) and set the related S3 Bucket to be used as source
 7. Define an environment variable named `DEST_BUCKET` and set its value to the name of the destination S3 bucket
 8. Upload some files to the source bucket
 9. Enjoy!

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/luisrpp/copy_to_another_s3_bucket.
