# AI Subdomain Generator - aisubs

This is a script that generates new subdomains based on a given subdomain using the power of the OpenAI GPT-3.5 Turbo model. It takes input from either a text file or command line and generates similar subdomains by interacting with the OpenAI API.

## Prerequisites

- Go programming language (version 1.16 or higher)
- OpenAI API key

## Installation

```bash
go install github.com/topscoder/aisubs@latest
```

## Usage

1. The script can be used in two ways:

    ```bash
    cat subdomains.txt | aisubs --apikey <OpenAI API Key> --amount 5
    ```

    Replace `<OpenAI API Key>` with your actual OpenAI API key.

2. Providing input via command line:

    ```bash
    echo "www.domain.com" | go run main.go --apikey <OpenAI API Key> --amount 5
    ```

    Replace `<OpenAI API Key>` with your actual OpenAI API key.

And it's able to use in your leet command chains. Eg.

```bash
 cat subdomains.txt | aisubs --apikey <OpenAI API Key> --amount 5 | httpx -ip -sc -cl -title -silent
```

The script reads input from either a text file or command line, where each line represents a subdomain in the format `<subdomain>.<domain>`. It generates `amount` number of new subdomains similar to the provided subdomain and prints them to the standard output.

## Configuration

Before running the script, make sure to obtain an API key by signing up on the [OpenAI website](https://openai.com/).

## Example

Let's say we have a text file called subdomains.txt with the following content:

```
www.example.com
api.example.com
```

We can generate 3 new subdomains similar to www.example.com using the script:

```bash
cat subdomains.txt | go run main.go --apikey <OpenAI API Key> --amount 3
```

The script will output something like this:

```
web.example.com
blog.example.com
shop.example.com
```

## Limitations

* The script relies on the OpenAI GPT-3.5 Turbo model and requires a valid OpenAI API key.
* The effectiveness of subdomain generation depends on the quality and diversity of the training data used to train the model.
* The script may encounter rate limits or errors when making API requests. It automatically handles rate limit errors by waiting for 20 seconds before retrying.
* The script assumes the input subdomains are in the format `<subdomain>.<domain>`.

## Contributing

If you would like to contribute to this project, please feel free to fork the repository and submit a pull request.


## License

This project is licensed under the MIT License.


## Contact

If you have any questions or feedback, please feel free to contact the project maintainers.
