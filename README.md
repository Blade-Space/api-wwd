# wwd - working with dirs 📃

Package for working with dirs

## import to progect

```golang
package main

import (
  wwf "api/wwd/routes"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  api := r.Group("/api/wwd")
  wwf.RegisterRoutes(api)

  r.Run(":3000")
}
```
# API Endpoints

<table><thead><tr><th>Method</th><th>Endpoint</th><th>JSON Request Body</th><th>Description</th></tr></thead><tbody><tr><td>POST</td><td>/delete_dir</td><td><code>{ "path": "&lt;directory_path&gt;" }</code></td><td>Delete a directory</td></tr><tr><td>POST</td><td>/rename_dir</td><td><code>{ "old_path": "&lt;old_path&gt;", "new_path": "&lt;new_path&gt;" }</code></td><td>Rename a directory</td></tr><tr><td>POST</td><td>/create_dir</td><td><code>{ "path": "&lt;directory_path&gt;" }</code></td><td>Create a new directory</td></tr><tr><td>GET</td><td>/work_dir</td><td>N/A</td><td>Get the current working directory</td></tr><tr><td>GET</td><td>/dir</td><td><code>{ "path": "&lt;directory_path&gt;" }</code></td><td>List all files within a directory and its subdirectories</td></tr><tr><td>POST</td><td>/copy_dir</td><td><code>{ "source_path": "&lt;source_path&gt;", "dest_path": "&lt;destination_path&gt;" }</code></td><td>Copy a directory to a new location</td></tr><tr><td>POST</td><td>/move_dir</td><td><code>{ "source_path": "&lt;source_path&gt;", "dest_path": "&lt;destination_path&gt;" }</code></td><td>Move a directory to a new location</td></tr></tbody></table>
