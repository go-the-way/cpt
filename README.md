# cpt
A simple slider captcha implementation in Go

Features
---
- **Gzip compression embedded**

Functions
---
- **SetTokenExpiration** `Set token expiration`
- **SetTokenClearJobExecTick** `Set token clear job ticker durtion`
- **SetTokenDeviation** `Set token deviation in px`
- **SetTokenLength** `Set token generate length`
- **SetResLoaderDefaultOpts** `Set default res loader opts`
- **Generate** `Generate the image token in base64`
- **Verify** `Verify the image token & xs`
- **Delete** `Delete the image token`
- **WrappedGenerateHandlerFunc** `Get the wrapped generate handler func`
- **WrappedVerifyHandlerFunc** `Get the wrapped verify handler func`
- **Serve** `Serve embedded http server`
- **ServeRoute** `Serve embedded http server with custom routers`

Embedded Routers
---

- **GET: /cpt/generate** `The Generate function called`

**_response_**
<pre>
{
    "bg_image_base_64": "data:image/png;base64,i",
    "bc_image_base_64": "data:image/png;base64,i",
    "token": "E4kAu5A2gjoXY7CfCw"
}
</pre>

- **GET: /cpt/generate?html** `The Generate function called, rendering html`

- **GET: /cpt/verify** `The Verify function called`

**_query parameters_**
<pre>
?token=
?x=
</pre>

**_response error_**
<pre>
{"err":"captcha_err"}
</pre>

**_response verified_**
<pre>
{"verified":"ok"}
</pre>

TODO
---
- **resloader_file** `Load resources from local file system`
- **resloader_uri** `Load resources from url address`