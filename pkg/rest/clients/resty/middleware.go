package pkcresty

import "github.com/go-resty/resty/v2"

// AddHeaderMiddleware agrega un middleware para añadir headers a todas las solicitudes.
func AddHeaderMiddleware(c Client, key, value string) {
	c.GetClient().OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		req.SetHeader(key, value)
		return nil
	})
}

// AddAuthTokenMiddleware agrega un middleware para incluir el token de autorización en las solicitudes.
func AddAuthTokenMiddleware(c Client, token string) {
	c.GetClient().OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		req.SetHeader("Authorization", "Bearer "+token)
		return nil
	})
}

// AddLoggingMiddleware agrega un middleware para registrar solicitudes y respuestas utilizando un Logger.
func AddLoggingMiddleware(c Client, logger Logger) {
	c.GetClient().OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		logger.Info("Request:", req.Method, req.URL, "headers:", req.Header)
		return nil
	})

	c.GetClient().OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
		logger.Info("Response:", resp.Status(), "for", resp.Request.Method, resp.Request.URL)
		return nil
	})
}
