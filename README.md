source: clear
features:
  - go
nginx:
  root: public_html/public
  passenger:
    enabled: "on"
    app_start_command: env PORT=$PORT ./main