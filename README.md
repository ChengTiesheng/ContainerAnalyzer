# ContainerAnalyzer

## Examples
- Analyse image from In-Cloud
```
$ ./ContainerAnalyzer docker://busybox
Getting image attribution information...

Prompt: Layer0 is the upper layer and layer1 is the bottom layer

============Layer1================
Type:      Docker Image Layer
Layer:     5c5fb281b01ee091a0fffa5b4a4c7fb7d358e7fb7c49c263d6d7a4e35d199fd0
Name:      registry-1.docker.io/library/busybox-5c5fb281b01ee091a0fffa5b4a4c7fb7d358e7fb7c49c263d6d7a4e35d199fd0
Tag:       latest
Version:   1.8.3
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-12-08T18:31:50Z
Comment:   
Parent:    
Checksum:  
App:

============Layer0================
Type:      Docker Image Layer
Layer:     fc0db02f30724abc777d7ae2b2404c6d074f1e2ceca19912352aea30a42f50b7
Name:      registry-1.docker.io/library/busybox-fc0db02f30724abc777d7ae2b2404c6d074f1e2ceca19912352aea30a42f50b7
Tag:       latest
Version:   1.8.3
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-12-08T18:31:51Z
Comment:   
Parent:    registry-1.docker.io/library/busybox-5c5fb281b01ee091a0fffa5b4a4c7fb7d358e7fb7c49c263d6d7a4e35d199fd0
Checksum:  
App:
	Exec:
		arg: /bin/sh
		arg: -c
		arg: "sh"
	User: 0
	Group: 0

```

- Analyse image file On-Premises

```
$ docker save -o redis.tar redis
$ ./ContainerAnalyzer redis.tar
Getting image attribution information...

Prompt: Layer0 is the upper layer and layer16 is the bottom layer

============Layer16================
Type:      dockerimage
Layer:     b014c4494ea56a76f945ba3d69b7d42c4a3031aef94f3ab6a53b506e354829bf
Name:      /redis-b014c4494ea56a76f945ba3d69b7d42c4a3031aef94f3ab6a53b506e354829bf
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T00:23:32Z
Comment:   
Parent:    
App:

============Layer15================
Type:      dockerimage
Layer:     2f329595e406d1adadb7e84bee918b9e495d3ebb9bae436f59652d6738dd3175
Name:      /redis-2f329595e406d1adadb7e84bee918b9e495d3ebb9bae436f59652d6738dd3175
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T00:23:34Z
Comment:   
Parent:    /redis-b014c4494ea56a76f945ba3d69b7d42c4a3031aef94f3ab6a53b506e354829bf
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0

============Layer14================
Type:      dockerimage
Layer:     2d34e9ec4a65e6eafc102c8409deb660f936aaef000c08e59be02e948c9c87f6
Name:      /redis-2d34e9ec4a65e6eafc102c8409deb660f936aaef000c08e59be02e948c9c87f6
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T10:58:32Z
Comment:   
Parent:    /redis-2f329595e406d1adadb7e84bee918b9e495d3ebb9bae436f59652d6738dd3175
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

============Layer13================
Type:      dockerimage
Layer:     8fc6e86792e468880a03554d71af4f4295b545715231d1656b08c43904e70526
Name:      /redis-8fc6e86792e468880a03554d71af4f4295b545715231d1656b08c43904e70526
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T10:58:49Z
Comment:   
Parent:    /redis-2d34e9ec4a65e6eafc102c8409deb660f936aaef000c08e59be02e948c9c87f6
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

============Layer12================
Type:      dockerimage
Layer:     867386488f3132428b94a0c1ceb2dbf4edce7332143a6d1bbd86c2cef00f938a
Name:      /redis-867386488f3132428b94a0c1ceb2dbf4edce7332143a6d1bbd86c2cef00f938a
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T10:58:52Z
Comment:   
Parent:    /redis-8fc6e86792e468880a03554d71af4f4295b545715231d1656b08c43904e70526
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

============Layer11================
Type:      dockerimage
Layer:     5e4fbefc8043207731f6841b9c80edb4f350f4986274a2b222c23c66e31d0567
Name:      /redis-5e4fbefc8043207731f6841b9c80edb4f350f4986274a2b222c23c66e31d0567
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T10:58:57Z
Comment:   
Parent:    /redis-867386488f3132428b94a0c1ceb2dbf4edce7332143a6d1bbd86c2cef00f938a
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

============Layer10================
Type:      dockerimage
Layer:     3894d23d17693c9b9f56ce23f3900c18c5166eff1ab4f1bd87ca3d9370e64c45
Name:      /redis-3894d23d17693c9b9f56ce23f3900c18c5166eff1ab4f1bd87ca3d9370e64c45
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:07:57Z
Comment:   
Parent:    /redis-5e4fbefc8043207731f6841b9c80edb4f350f4986274a2b222c23c66e31d0567
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"

============Layer9================
Type:      dockerimage
Layer:     552c9d8714969f10d8810b586dd0b2bfa77cdd34a7e75610b43cd2437d6bc9c5
Name:      /redis-552c9d8714969f10d8810b586dd0b2bfa77cdd34a7e75610b43cd2437d6bc9c5
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:07:58Z
Comment:   
Parent:    /redis-3894d23d17693c9b9f56ce23f3900c18c5166eff1ab4f1bd87ca3d9370e64c45
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"

============Layer8================
Type:      dockerimage
Layer:     1585b92a2e813a40bb86e3919d9765b15ac8d3e2cf363794a917f27d9914756b
Name:      /redis-1585b92a2e813a40bb86e3919d9765b15ac8d3e2cf363794a917f27d9914756b
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:07:58Z
Comment:   
Parent:    /redis-552c9d8714969f10d8810b586dd0b2bfa77cdd34a7e75610b43cd2437d6bc9c5
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"

============Layer7================
Type:      dockerimage
Layer:     f52677abfb05adeaac2fc8d820b86f37a895bf38ce042fdf4322a057ce7443f0
Name:      /redis-f52677abfb05adeaac2fc8d820b86f37a895bf38ce042fdf4322a057ce7443f0
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:49Z
Comment:   
Parent:    /redis-1585b92a2e813a40bb86e3919d9765b15ac8d3e2cf363794a917f27d9914756b
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"

============Layer6================
Type:      dockerimage
Layer:     5228fb0bad0a9df970d11d0ea2614daa8c85d270029d50aab3d5e0399ad4f789
Name:      /redis-5228fb0bad0a9df970d11d0ea2614daa8c85d270029d50aab3d5e0399ad4f789
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:50Z
Comment:   
Parent:    /redis-f52677abfb05adeaac2fc8d820b86f37a895bf38ce042fdf4322a057ce7443f0
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"

============Layer5================
Type:      dockerimage
Layer:     c7e038aa4e474058065109a731092dd80fa54daf8d37c825065e6403c6ab4d22
Name:      /redis-c7e038aa4e474058065109a731092dd80fa54daf8d37c825065e6403c6ab4d22
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:51Z
Comment:   
Parent:    /redis-5228fb0bad0a9df970d11d0ea2614daa8c85d270029d50aab3d5e0399ad4f789
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false

============Layer4================
Type:      dockerimage
Layer:     179d219ba2154d2bc110e376e0c3924f9380d2920f3cc01f2b94013c2ec36618
Name:      /redis-179d219ba2154d2bc110e376e0c3924f9380d2920f3cc01f2b94013c2ec36618
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:51Z
Comment:   
Parent:    /redis-c7e038aa4e474058065109a731092dd80fa54daf8d37c825065e6403c6ab4d22
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	WorkingDirectory: /data
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false

============Layer3================
Type:      dockerimage
Layer:     4a2fca62f85ed69768a06432b2ce951aab8353686d0341e7761acbd7b381667b
Name:      /redis-4a2fca62f85ed69768a06432b2ce951aab8353686d0341e7761acbd7b381667b
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:52Z
Comment:   
Parent:    /redis-179d219ba2154d2bc110e376e0c3924f9380d2920f3cc01f2b94013c2ec36618
App:
	Exec:
		arg: /bin/bash
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	WorkingDirectory: /data
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false

============Layer2================
Type:      dockerimage
Layer:     57dc1fc95252b45059dadc6ce5fb1d75f2ca62c3aef8912aa1a1964262b22d1f
Name:      /redis-57dc1fc95252b45059dadc6ce5fb1d75f2ca62c3aef8912aa1a1964262b22d1f
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:52Z
Comment:   
Parent:    /redis-4a2fca62f85ed69768a06432b2ce951aab8353686d0341e7761acbd7b381667b
App:
	Exec:
		arg: /entrypoint.sh
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	WorkingDirectory: /data
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false

============Layer1================
Type:      dockerimage
Layer:     a8b4e95f4daea4c45f2e38427c97844db9e6d40d0444056ff7fe012b5151c1e7
Name:      /redis-a8b4e95f4daea4c45f2e38427c97844db9e6d40d0444056ff7fe012b5151c1e7
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:53Z
Comment:   
Parent:    /redis-57dc1fc95252b45059dadc6ce5fb1d75f2ca62c3aef8912aa1a1964262b22d1f
App:
	Exec:
		arg: /entrypoint.sh
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	WorkingDirectory: /data
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false
	Ports:
		name: "6379/tcp", protocol: "tcp", port: 6379

============Layer0================
Type:      dockerimage
Layer:     82ca9f96ee0582b7f95aefb9872db606df4cf7254e03caa19a5f16ab8dfb6052
Name:      /redis-82ca9f96ee0582b7f95aefb9872db606df4cf7254e03caa19a5f16ab8dfb6052
Version:   latest
OS:        linux
Arch:      amd64
Author:    
Epoch:     2015-11-20T11:08:53Z
Comment:   
Parent:    /redis-a8b4e95f4daea4c45f2e38427c97844db9e6d40d0444056ff7fe012b5151c1e7
App:
	Exec:
		arg: /entrypoint.sh
		arg: redis-server
	User: 0
	Group: 0
	Environment:
		name :"PATH", value :"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		name :"REDIS_VERSION", value :"3.0.5"
		name :"REDIS_DOWNLOAD_URL", value :"http://download.redis.io/releases/redis-3.0.5.tar.gz"
		name :"REDIS_DOWNLOAD_SHA1", value :"ad3ee178c42bfcfd310c72bbddffbbe35db9b4a6"
	WorkingDirectory: /data
	MountPoints:
		name: "volume/data", path: "/data", readOnly: false
	Ports:
		name: "6379/tcp", protocol: "tcp", port: 6379

```
