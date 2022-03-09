# DCT Go Client SDK

This is a Go Client SDK leveraging the DCT APIGW.

For SDK usage, please refer to [OPENAPI-README.md](OPENAPI-README.md) file

# DCT Go Client SDK Generator Script

This is a guide on generation of the client SDK in Go Lang for the DCT Orbital APIGW.

## Getting Started

The reources for code generation are present in the 'generation-scripts' directory.
The script codegen.sh generates a GO Client SDK using the openapi-generator-cli tool. It generates the SDK, initializes the SDK and runs a test go function.

### Prerequisites

Java needs to be installed to run the jar file. And Golang >= 1.17 needs to be installed in the machine where the script is run.

Refer to the official docs of the respective languages for guidance.

Also the script requires an API Spec file for the DCT-APIGW. A sample file is present in the directory. For the latest spec, please download it from https://raw.githubusercontent.com/delphix/orbital-api-gateway/main/app/src/main/resources/api.yaml

### Running the Script

As mentioned above the script requires a spec file.

Run the script as follows

```
sh codegen.sh api.yaml
```

On successful run, following tasks will result:

```
1. Generation of client SDK
2. Initialize the SDK module
3. Run the Test Go file, input the API key and hostname of the Delphix Engine and display list of Engines added to the APIGW
```

## Running the tests

The test.go file executes a get all engines operation on the configured DCT-APIGW. 
To run the test on your delphix engine, provide the API key and Hostname of the engine when prompted.

### Break down into end to end tests

The test essentially tests if the client SDK is functional and should be run before raising a PR.
The result of the test will be available at the end of the script as follows.

```
List of engines are:
my-host.domain.co
```

If in any case, the test fails, we will be getting a 'test failed' message in the end.

## Deployment

Once successful generation of client SDK and the test succeeds, we are ready to raise a PR.

## Contributing

Please read [CONTRIBUTING.md](https://github.com/delphix/.github/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- **Uddipaan Hazarika** - _Initial work_ - [Company](https://github.com/delphix/)

## License

This project is licensed under the XX License - see the [LICENSE.md](LICENSE.md) file for details
