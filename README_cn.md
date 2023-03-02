<div>
<a href="https://alist.nn.ci"><img height="100px" alt="logo" src="https://cdn.jsdelivr.net/gh/alist-org/logo@main/logo.svg"/></a>
  <p><em>🗂一个基于开源项目 AList 修改，符合个人使用需求的云盘项目。</em></p>
  <a href="https://github.com/Xhofe/alist/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Xhofe/alist" alt="License" />
  </a>
</div>

---

## 功能

- [x] 多种存储
    - [x] 本地存储
    - [x] [阿里云盘](https://www.aliyundrive.com/)
    - [x] [百度网盘](http://pan.baidu.com/)
    - [x] OneDrive / Sharepoint（[国际版](https://www.office.com/), [世纪互联](https://portal.partner.microsoftonline.cn),de,us）
    - [x] FTP / SFTP
    - [x] WebDav(支持无API的OneDrive/SharePoint)
- [x] 部署方便，开箱即用
- [x] 文件预览（PDF、markdown、代码、纯文本……）
- [x] 画廊模式下的图像预览
- [x] 视频和音频预览，支持歌词和字幕
- [x] Office 文档预览（docx、pptx、xlsx、...）
- [x] `README.md` 预览渲染
- [x] 文件永久链接复制和直接文件下载
- [x] 黑暗模式
- [x] 国际化
- [x] 受保护的路由（密码保护和身份验证）
- [x] WebDav (具体见 https://alist.nn.ci/zh/guide/webdav.html)
- [x] [Docker 部署](https://hub.docker.com/r/xhofe/alist)
- [x] Cloudflare workers 中转
- [x] 文件/文件夹打包下载
- [x] 网页上传(可以允许访客上传)，删除，新建文件夹，重命名，移动，复制
- [x] 离线下载
- [x] 跨存储复制文件

## 许可

`AList` 是在 AGPL-3.0 许可下许可的开源软件。

## 免责声明
- 本程序为免费开源项目，只是方便个人使用下载以及学习golang，使用时请遵守相关法律法规，请勿滥用；
- 本程序通过调用官方sdk/接口实现，无破坏官方接口行为；
- 本程序仅做302重定向/流量转发，不拦截、存储、篡改任何用户数据.