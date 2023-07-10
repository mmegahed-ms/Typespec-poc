# Standardize API Design and Development: Automating the End-to-End Workflow with Typespec, Kiota, Azure API Management, and GitHub Actions

## Build Status

| GitHub Action | Status |
| ----------- | ----------- |
| Build | [![Build](https://github.com/mmegahed-ms/Typespec-poc/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/mmegahed-ms/Typespec-poc/actions/workflows/build.yml) |
| Release | [![Release](https://github.com/mmegahed-ms/Typespec-poc/actions/workflows/release.yml/badge.svg)](https://github.com/mmegahed-ms/Typespec-poc/actions/workflows/release.yml) |

## About

In today's fast-paced software development landscape, automating the API design and development process can greatly enhance productivity and efficiency. This article explores how to streamline the entire workflow using a comprehensive toolset that includes Typespec, Kiota, Azure API Management, and GitHub Actions.

We start by introducing Typespec,TypeSpec is a language for describing cloud service APIs and generating other API description languages, client and service code, documentation, and other assets. TypeSpec provides highly extensible core language primitives that can describe API shapes common among REST, GraphQL, gRPC, and other protocols.

Next, we delve into Kiota, a command line tool for generating an API client to call any OpenAPI-described API you are interested in. The goal is to eliminate the need to take a dependency on a different API SDK for every API that you need to call. Kiota API clients provide a strongly typed experience with all the features you expect from a high quality API SDK, but without having to learn a new library for every HTTP API.


Additionally, we explore the capabilities of Azure API Management, a comprehensive service that simplifies the management and publication of APIs. Azure API Management allows you to seamlessly import your OpenAPI specifications, set up mock services for testing, and expose your APIs to consumers securely. Furthermore, we highlight how GitHub Actions can be utilized to automate the entire process. With GitHub Actions, you can define workflows that automatically compile Typespec into OpenAPI specifications, import them into Azure API Management, and generate SDKs using Kiota.

By combining Typespec, Kiota, Azure API Management, and GitHub Actions, you can achieve a fully automated end-to-end workflow for API design and development. This automation not only improves consistency and reduces development time but also enhances collaboration among teams by providing a centralized and efficient platform.

Whether you are a developer, architect, or product owner, this article equips you with the knowledge and insights to embark on a journey towards automating and standardizing API design and development. Discover how these powerful tools can transform your API workflow, saving you time, effort, and resources in the process.

## Architecture

![image](https://github.com/mmegahed-ms/Typespec-poc/blob/ca7e5ace6de5bb892d7b860002bfcb0af40d2287/docs/images/Typspec-kiota-arch.png)

## Prerequisites

* Install [Visual Studio Code](https://code.visualstudio.com/download)
* Install [Typespec](https://microsoft.github.io/typespec/introduction/installation) Extension for Visual Studio Code.
* Install [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) Extension for Visual Studio Code.
* Install [OpenAPI Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) Extension for Visual Studio Code.
* Install Chocolatey (package manager)

```ps1
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

* Install Bicep CLI

```ps1
choco install bicep
```

* Install Az Module in PowerShell

```ps1
Install-Module -Name Az -AllowClobber -Scope CurrentUser
```

## Create an OpenAPI Specification with Typespec

In Visual Studio Code, start creating a new Typespec project...

```ps1
tsp init
```
This will prompt you with a few question, pick the Generic Rest API template, your project name, and select the @typespec/openapi3 library.

Next, you can install the dependencies

```ps1
tsp install
```

```ps1
You should now have a basic TypeSpec project setup with a structure looking like

package.json      # Package manifest defining your typespec project as a node package.
tspconfig.yaml # TypeSpec project configuration letting you configure emitters, emitter options, compiler options, etc.
main.tsp         # TypeSpec entrypoint
```
you then can define your API in the main.tsp file 
 
I started with a sample json customer object to build an inteface for this repo 
(https://github.com/mmegahed-ms/Typespec-poc/blob/5639d46603e780740633b7b0d47c3c9d9251f801/sample.json)


* Delete the backend server reference from the openapi specification (when generated)

```json
"servers": [
        {
            "url": "http://localhost:5000"
        }
    ],
```

## Implement a mocking response policy

You can add a mocking response easily in a policy (see snippet below). I've added this policy in the automated deployment [here](deploy/release/policies/api_policy.xml), which will return a mocking response.

```xml
<mock-response status-code="200" content-type="application/json" />
```

## Deploy Manually

* Git Clone the repository

```ps1
git clone https://github.com/pascalvanderheiden/ais-apim-copilot.git
```

* Deploy it all by one script

I've included all the steps in 1 Powershell script. This will create all the needed resources. Keep in mind that this will take a while to deploy.

I've used these variables:

```ps1
$subscriptionId = "<subscription_id>"
$namePrefix = "<project_prefix>"

# For removing soft-delete
$apimName = "<apim_name>"
```

```ps1
.\deploy\manual-deploy.ps1 -subscriptionId $subscriptionId -namePrefix $namePrefix
```

* Remove the APIM Soft-delete

If you deleted the deployment via the Azure Portal, and you want to run this deployment again, you might run into the issue that the APIM name is still reserved because of the soft-delete feature. You can remove the soft-delete by using this script:

```ps1
.\deploy\del-soft-delete-apim.ps1 -subscriptionId $subscriptionId -apimName $apimName
```

* Testing

I've included a tests.http file with relevant Test you can perform, to check if your deployment is successful. Or you can just test it via the Azure Portal.

![ais-apim-copilot](docs/images/apim_result.png)

## Deploy with Github Actions

* Fork this repository

* Generate a Service Principal

```ps1
az ad sp create-for-rbac -n <name_sp> --role Contributor --sdk-auth --scopes /subscriptions/<subscription_id>
```

Copy the json output of this command.

* Update GitHub Secrets for customizing your deployment

In the repository go to 'Settings', on the left 'Secrets', 'Actions'.
And pass the json output in the command used above into the secret 'AZURE_CREDENTIALS'.

The following secrets need to be created:

* AZURE_CREDENTIALS
* AZURE_SUBSCRIPTION_ID
* LOCATION
* PREFIX

### Commit

Commit the changes, and this will trigger the CI Build Pipeline.
