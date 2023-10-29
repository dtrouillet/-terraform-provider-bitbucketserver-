# Data Source: bitbucketserver_project

This data source allows you to retrieve project details.

## Example Usage

```hcl
data "bitbucketserver_project" "my-project" {
  key = "MYPROJECT"
}
```

## Argument Reference

* `key` - Unique key of the project.

## Attribute Reference

* `name` - Name of the project.
* `description` - Description of the project. 
* `public` - Set to `true` if the project is public. 
* `avatar` - Avatar url of the project.
* `repos` - List of repos slugs of the project.
