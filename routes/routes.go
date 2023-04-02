package wwd

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.POST("/delete_dir", DeleteDir)
	api.POST("/rename_dir", RenameDir)
	api.POST("/create_dir", CreateDir)
	api.GET("/work_dir", TakeWorkDir)
	api.GET("/dir", GetDirTree)
	api.POST("/copy_dir", CopyDir)
	api.POST("/move_dir", MoveDir)
}

type DirRequest struct {
	Path string `json:"path"`
}

type RenameRequest struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

type CopyMoveRequest struct {
	SourcePath string `json:"source_path"`
	DestPath   string `json:"dest_path"`
}

// * Удаление директории
func DeleteDir(c *gin.Context) {
	var req DirRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = os.RemoveAll(req.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

// * Изменить имя директории
func RenameDir(c *gin.Context) {
	var req RenameRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = os.Rename(req.OldPath, req.NewPath)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

// *  Создать новую папку
func CreateDir(c *gin.Context) {
	var req DirRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = os.MkdirAll(req.Path, 0755)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

// * Возвращает рабочую директорию
func TakeWorkDir(c *gin.Context) {
	wd, err := os.Getwd()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"work_dir": wd})
}

// * Получение всех файлов папки (если тоже папки то и их содержимое)
func GetDirTree(c *gin.Context) {
	var req DirRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	var files []string
	err = filepath.Walk(req.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"files": files})
}

// * Скопировать папку в указанный путь
func CopyDir(c *gin.Context) {
	var req CopyMoveRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = filepath.Walk(req.SourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(req.SourcePath, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(req.DestPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		return err
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

// * Переместить папку в указанный путь
func MoveDir(c *gin.Context) {
	var req CopyMoveRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = os.Rename(req.SourcePath, req.DestPath)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}
