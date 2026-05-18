package scp

import (
	tea "github.com/charmbracelet/bubbletea"
	internalscp "github.com/mheaton92/quay/internal/scp"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch m.step {

	case StepDirection:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.String() {
			case "u":
				m.upload = true
				m.step = StepFilePicker
				return m, m.fp.Init()
			case "d":
				m.upload = false
				m.step = StepRemotePath
			}
		}

	case StepFilePicker:
		var cmd tea.Cmd
		m.fp, cmd = m.fp.Update(msg)
		if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
			if m.upload {
				m.localPath = path
				m.step = StepRemotePath
			} else {
				m.localPath = path
				m.step = StepTransferring
				err := internalscp.Download(m.conn, m.remotePath, m.localPath)
				if err != nil {
					m.err = err
					m.status = "Download failed: " + err.Error()
				} else {
					m.status = "Download complete!"
				}
			}
		}
		return m, cmd

	case StepRemotePath:
		var cmd tea.Cmd
		m.remoteInput, cmd = m.remoteInput.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok {
			if key.String() == "enter" {
				m.remotePath = m.remoteInput.Value()
				if m.upload {
					m.step = StepTransferring
					err := internalscp.Upload(m.conn, m.localPath, m.remotePath)
					if err != nil {
						m.err = err
						m.status = "Upload failed: " + err.Error()
					} else {
						m.status = "Upload complete!"
					}
				} else {
					m.step = StepLocalPath
				}
				return m, nil
			}
		}
		return m, cmd

	case StepTransferring:
		// waiting for user to press esc

	case StepLocalPath:
		var cmd tea.Cmd
		m.localInput, cmd = m.localInput.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok {
			if key.String() == "enter" {
				m.localPath = m.localInput.Value()
				m.step = StepTransferring
				err := internalscp.Download(m.conn, m.remotePath, m.localPath)
				if err != nil {
					m.err = err
					m.status = "Download failed: " + err.Error()
				} else {
					m.status = "Download complete!"
				}
				return m, nil
			}
		}
		return m, cmd
	}

	return m, nil
}
