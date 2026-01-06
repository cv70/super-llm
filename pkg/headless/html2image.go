package headless

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// RemoteHTML2Image 连接远程浏览器容器将HTML渲染为图片
func RemoteHTML2Image(remoteURL, htmlContent string) ([]byte, error) {
	// 连接远程浏览器
	browser := rod.New().ControlURL(remoteURL).MustConnect()
	defer browser.MustClose()

	return renderHTMLToImage(browser, htmlContent)
}

// renderHTMLToImage 通用的HTML渲染为图片函数
func renderHTMLToImage(browser *rod.Browser, htmlContent string) ([]byte, error) {
	htmlContent = "<div id='selector'>" + htmlContent + "</div>"
	
	// 打开空白页
	page, err := browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		return nil, err
	}
	
	// 设置页面内容（注入HTML）
	err = page.SetDocumentContent(htmlContent)
	if err != nil {
		return nil, err
	}
	
	// 等待页面加载完成
	err = page.WaitStable(time.Second)
	if err != nil {
		return nil, err
	}
	
	// 等待目标元素出现
	element, err := page.Element("#selector")
	if err != nil {
		return nil, err
	}
	
	// 返回截图
	return element.Screenshot(proto.PageCaptureScreenshotFormatPng, 100)
}
