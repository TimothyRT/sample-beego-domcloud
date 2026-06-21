source: clear
features:
  - go
nginx:
  root: public_html/public
  passenger:
    enabled: "on"
    startup_file: app