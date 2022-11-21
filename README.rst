=====================================
xkcd password generator for Terraform
=====================================

You know: https://xkcd.com/936/

Wraps https://github.com/martinhoefling/goxkcdpwgen so credit there.

Usage
-----

See examples folder::

    terraform apply && terraform show -json | jq .values.outputs

Development
-----------

Remember to set your overriders in `.terraformrc`, see https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers

- change to a branch::

    git checkout -b my_branch

- Install project deps and pre-commit hooks::

    pre-commit install
    pre-commit run --all-files

- Ready to go.
