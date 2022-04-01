# remote-config

This is a lightweight and simple remote config server with a UI for management backed by AWS S3.

The backend API (Go) serves a frontend (Vue) and HTTP endpoints.

![demo](https://user-images.githubusercontent.com/10101283/143332259-3b12b634-b15d-4568-bf24-c3d1d8173a0e.png)

> I wrote this because I don't want to pay for another service and Firebase doesn't work all the time in web extensions :/

## Get Started

1. Create S3 Bucket
1. Build the docker image
1. Deploy the docker image

That's it!

### 1. Create S3 Bucket

You can create the S3 bucket however you like. Or you can use an existing one.

See AWS's docs if you've never done this before: <https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html>

### 2. Build the docker image

Eventually, I'm going to publish this image to docker hub. Until then, build and tag the image yourself:

```bash
git clone git@github.com:anime-skip/remote-config.git
docker build . -t my-remove-config
```

### 3. Deploy the docker image

You will have to self host this docker image. AWS, Heroku, GoogleCloud, anywhere that supports deploying docker images will work.

Make sure you provide all the [environment variables](#environment-variables) however your hosting provider allows.

---

## Docs

### Environment Variables

| Variable              | Required | Default | Example                                         | Description                                                                                                                                                                                       |
| :-------------------- | :------: | :-----: | :---------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| PORT                  |    ✅    |   ---   | `80`, `8080`                                    | The port to run the application on                                                                                                                                                                |
| AUTH_TOKEN            |    ✅    |   ---   | `password`                                      | The token used to authenticate changes to the remove config. Config is viewable by everyone, but only writable if you know the token                                                              |
| AWS_REGION            |    ✅    |   ---   | `us-east-1`                                     | The AWS region your bucket is in                                                                                                                                                                  |
| AWS_ACCESS_KEY_ID     |    ✅    |   ---   |                                                 | See [AWS Docs](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys)                                                                                                |
| AWS_SECRET_ACCESS_KEY |    ✅    |   ---   |                                                 | See [AWS Docs](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys)                                                                                                |
| S3_BUCKET             |    ✅    |   ---   | `remote-config`                                 | The name of the bucket used as storage                                                                                                                                                            |
| S3_FILE_PATH          |    ✅    |   ---   | `dev.json`, `prod.json`, `path/to/my-file.json` | The path to the file in your bucket. The file can be named whatever, as long as it ends with `.json`, and the path can excluded or point to any subdirectory. Missing directories will be created |

### API Docs

[HTTPie](https://httpie.io/) is used for example calls because it's easier to read and understand compared to curl.

> `http -b` prints just the response body.

#### `GET /api/apps`

List all the apps you have created remote config for.

```bash
$ http -b "remote-config.anime-skip.com/api/apps"
[
    "Anime Skip Player",
    "anime-skip.com"
]
```

#### `GET /api/config/{app}`

Get the current config for an app.

```bash
$ http -b "remote-config.anime-skip.com/api/config/anime-skip.com"
{
    "array": [
        1,
        null,
        "string"
    ],
    "number": 69,
    "object": {
        "key": "value"
    },
    "optional": null,
    "string": "test"
}
```

#### `PUT /api/config/{app}`

Create or update the config for a given app name, returning the updated config.

**Requires `Authorization: Bearer <token>` header**

```bash
$ http -b -A bearer -a "<token>" PUT "remote-config.anime-skip.com/api/config/example" key=value
{
    "key": "value"
}
```

#### `DELETE /api/config/{app}`

Delete an app's config, returning the deleted config

**Requires `Authorization: Bearer <token>` header**

```bash
$ http -b -A bearer -a "<token>" DELETE "remote-config.anime-skip.com/api/config/example"
{
    "key": "value"
}
```

---

## Other Tools

- JS client with caching and a synchronous API, based off Firebase: <https://github.com/anime-skip/remote-config-client-ts>
