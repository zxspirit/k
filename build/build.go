package build

import "time"

// Server 函数选项模式
type Server struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}
type Option func(*Server)

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.Addr = addr
	}

}
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = readTimeout
	}
}
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = writeTimeout
	}
}
func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func NewServer(opts ...Option) *Server {
	s := &Server{
		Addr:         "localhost",
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Timeout:      10 * time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
