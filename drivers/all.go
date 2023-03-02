package drivers

import (
	_ "github.com/alist-org/alist/v3/drivers/aliyundrive_open"
	_ "github.com/alist-org/alist/v3/drivers/baidu_netdisk"
	_ "github.com/alist-org/alist/v3/drivers/ftp"
	_ "github.com/alist-org/alist/v3/drivers/local"
	_ "github.com/alist-org/alist/v3/drivers/onedrive"
	_ "github.com/alist-org/alist/v3/drivers/sftp"
)

// All do nothing,just for import
// same as _ import
func All() {

}
