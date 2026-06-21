```yaml
source: clear
features:
  - go
nginx:
  root: public_html/public
  passenger:
    enabled: on
    app_root: /home/witty-zef/public_html/public
    app_start_command: ./app
    app_type: generic
```

```nginx
server {
    server_name www.witty-zef.domcloud.dev witty-zef.domcloud.dev;
    listen 152.53.38.7:443 ssl;
    listen [2a0a:4cc0:2000:332::]:443 ssl;
    root /home/witty-zef/public_html/public;
    index index.html index.php;
    access_log /var/log/virtualmin/witty-zef.domcloud.dev_access_log;
    error_log /var/log/virtualmin/witty-zef.domcloud.dev_error_log;
    location /static/ {
        root /home/witty-zef/ego/app/;
        expires 30d;
    }
    location / {
        proxy_pass http://127.0.0.1:3000;
    }
    location ~ \.php$ {
        return 404;
        fastcgi_pass unix:/run/php-fpm/17820541692399210.sock;
    }
    ssl_certificate /home/domcloud/ssl.combined;
    ssl_certificate_key /home/domcloud/ssl.key;
    passenger_enabled on;
    passenger_app_root "/home/witty-zef/public_html/public";
    passenger_app_start_command "./app";
    passenger_app_type generic;
}
```