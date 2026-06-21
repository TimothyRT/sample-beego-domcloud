This sample Beego project is fitted with a custom GitHub Actions workflow to build the software and deploy it to a DOM Cloud instance.

## Deploying to DOM Cloud

### Boot recipe

```yaml
source: clear
features:
  - go
nginx:
  root: public_html/public
  passenger:
    enabled: on
    app_root: /public_html/public
    app_start_command: env PORT=$PORT ./app
    app_type: generic
```

### Building

The 'Build to DOM Cloud' workflow (`domcloud.yaml`) builds and deploys the project files to a DOM Cloud instance. This workflow is run automatically when a commit is pushed to GitHub. You can also opt to run it manually if needed. Regardless, make sure to setup an SSH public/private key pair first!
