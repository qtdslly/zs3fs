zs3fs allows you to mount an S3 bucket as a filey system.

It's a Filey System instead of a File System because zs3fs strives
for performance first and POSIX second. Particularly things that are
difficult to support on S3 or would translate into more than one
round-trip would either fail (random writes) or faked (no per-file
permission). 

# Usage

```ShellSession
$ cat ~/.aws/credentials
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
$ $GOPATH/bin/zs3fs <bucket> <mountpoint>
$ $GOPATH/bin/zs3fs <bucket:prefix> <mountpoint> # if you only want to mount objects under a prefix
```

Users can also configure credentials via the
[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
or the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.

To mount an S3 bucket on startup, make sure the credential is
configured for `root`, and can add this to `/etc/fstab`:

```
zs3fs#bucket   /mnt/mountpoint        fuse     _netdev,allow_other,--file-mode=0666,--dir-mode=0777    0       0
```

See also: [Instruction for Azure Blob Storage, Azure Data Lake Gen1, and Azure Data Lake Gen2](https://zs3fs/blob/master/README-azure.md).

Got more questions? Check out [questions other people asked](https://zs3fs/issues?utf8=%E2%9C%93&q=is%3Aissue%20label%3Aquestion%20)