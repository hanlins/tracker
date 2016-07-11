##DEPLOYMENT

Nginx is used. Copy configuration file to /etc/nginx/sites-available and then
cd to sites-enabled then link (ln -s source target) that file. Also notice that
tester should change his/her local /etc/hosts to visit the trackers' domain.
(i.e. tracker[a-e].com).

Thanks to Roesner's paper (`Detecting and defending against third-party tracking
 on the web`), trackers we are going to try are implemented in 5 categories.

##DIRECTORY

`+-- etc      ==>` configurations

`|`

`+-- src      ==>` golang server

`|`

`+-- trackera ==>` resource for type - A tracker

`|`

 `...         ==>` remaining trackers
