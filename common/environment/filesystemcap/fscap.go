package filesystemcap

import "github.com/robovpn/github.com/robovpn/v2fly_core/common/platform/filesystem/fsifce"

type FileSystemCapabilitySet interface {
	OpenFileForReadSeek() fsifce.FileSeekerFunc
	OpenFileForRead() fsifce.FileReaderFunc
	OpenFileForWrite() fsifce.FileWriterFunc
}
