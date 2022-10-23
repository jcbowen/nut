package nut

import "errors"

// errorForMessage returns an error for the specified NUT error code.
func errorForMessage(message string) (err error) {
	switch message {
	case "ACCESS-DENIED":
		err = errors.New("客户端的主机和或身份验证详细信息（用户名、密码）不足以执行请求的命令。")
	case "UNKNOWN-UPS":
		err = errors.New("请求中指定的UPS未知。这通常意味着它与ups.conf中的任何内容都不匹配")
	case "VAR-NOT-SUPPORTED":
		err = errors.New("指定的UPS不支持请求中的变量。对于upsd（如服务器）处理的空间中的未识别变量，也会发送此消息。")
	case "CMD-NOT-SUPPORTED":
		err = errors.New("指定的UPS不支持请求中的即时命令")
	case "INVALID-ARGUMENT":
		err = errors.New("客户端向命令发送了一个参数，该参数在该上下文中无法识别或无效。这通常是由于使用无效子命令发送有效命令（如GET）导致的")
	case "INSTCMD-FAILED":
		err = errors.New("upsd无法将即时命令请求传递给驱动程序。客户无法获得更多信息。这通常表示驾驶员死亡或受伤")
	case "SET-FAILED":
		err = errors.New("upsd无法将设置请求传递给驱动程序。这就像上面的INSTCMD-FAILED")
	case "READONLY":
		err = errors.New("SET命令中请求的变量不可写")
	case "TOO-LONG":
		err = errors.New("SET命令中请求的值太长")
	case "FEATURE-NOT-SUPPORTED":
		err = errors.New("upsd的此实例不支持请求的功能。这目前仅用于TLS/SSL模式（STARTTLS）")
	case "FEATURE-NOT-CONFIGURED":
		err = errors.New("upsd的此实例未正确配置以允许请求的功能运行。目前这也仅限于STARTTLS")
	case "ALREADY-SSL-MODE":
		err = errors.New("此连接上已启用TLS/SSL模式，因此upsd无法再次启动")
	case "DRIVER-NOT-CONNECTED":
		err = errors.New("upsd无法执行请求的命令，因为该UPS的驱动程序未连接。这通常意味着驾驶员没有在运行，如果是，则是ups.conf配置错误")
	case "DATA-STALE":
		err = errors.New("upsd已连接到UPS的驱动程序，但该驱动程序未提供定期更新或已将数据明确标记为过时。upsd拒绝提供过时单位的变量，以避免错误读数。这通常意味着驱动程序正在运行，但它与硬件失去了通信。检查设备的物理连接")
	case "ALREADY-LOGGED-IN":
		err = errors.New("客户端已发送了UPS的登录信息，无法再次发送。目前，每个连接只能有一条登录记录")
	case "INVALID-PASSWORD":
		err = errors.New("客户端发送了无效的密码-可能是空密码")
	case "ALREADY-SET-PASSWORD":
		err = errors.New("客户端已设置密码，无法设置另一个密码。正常的NUT客户端也不应出现这种情况")
	case "INVALID-USERNAME":
		err = errors.New("客户端发送了无效的USERNAME")
	case "ALREADY-SET-USERNAME":
		err = errors.New("客户端已经设置了USERNAME，无法再设置一个。正常的NUT客户端不应出现这种情况")
	case "USERNAME-REQUIRED":
		err = errors.New("请求的命令需要用户名进行身份验证，但客户端尚未设置用户名")
	case "PASSWORD-REQUIRED":
		err = errors.New("请求的命令需要一个用于身份验证的密码，但客户端尚未设置")
	case "UNKNOWN-COMMAND":
		err = errors.New("upsd无法识别请求的命令")
	case "INVALID-VALUE":
		err = errors.New("请求中指定的值无效。这通常适用于ENUM类型的SET，该SET使用的值不在允许值列表中")
	default:
		err = errors.New("未知的错误码")
	}

	return err
}
