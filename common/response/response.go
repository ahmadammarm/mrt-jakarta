package response

type APIResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
}

