server {
	listen 80; 		# ipv4
	listen [::]:80; 	# ipv6

	server_name trackerA.com trackerB.com trackerC.com trackerD.com trackerE.com; 	# vhost

	location / {
		#root /root/mylab/;
		proxy_pass http://127.0.0.1:8002;
	}

	#location /static/ {
		
}
