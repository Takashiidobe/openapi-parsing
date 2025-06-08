# Notes for Generating All ARM clients

It's also possible to go straight from the swagger file to a go file
using autorest.

If so, I can parse the go files and generate clients for all of them.
It should also be possible to generate the YAML spec files if needed.

- Generate the openapi v3 file, where I can find out which list + get
  requests are supported + their return types for the response
- Next, generate the entire go directory for the resourceGroup using the
  swagger file
- Use asty to turn the go files into json, so we can use any language to
  parse the relevant information
- Now we just need to parse the openapi + go ast. The rest api specs
  should have the URL path to return types, like /storageAccounts to
  StorageAccountResponse which would store a []StorageAccount

- When you pass in a URL path to generate a particular crawler path, we
  need to merge them (with any possible child crawlers), so deduplicate
  by sorting and removing any shorter neighbors

## Information Required

- URL path
- Method (GETs only)
- Parameters (url params like resourceGroupName)
    - Special handling of params like top
- Return type of the 200 response -> mapped to Go code (*ListResponse)
- Then generate the resource name from the *ListResponse

- Also must generate intermediate ChannelSteps + ResourceSteps +
  PayloadSteps.
