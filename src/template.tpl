### SCDS Generated config ###

{{ range .Hosts }}
Host {{ range .Host }} {{ . }} {{ end }}
 Hostname {{ .Hostname }}
{{ if .Port }} Port {{ .Port}}{{ end }}
{{- if .User }} User {{ .User}}{{ end }}
{{- if .ProxyJump }} ProxyJump {{ .ProxyJump}}{{ end }}
{{- if .ProxyCommand }} ProxyCommand {{ .ProxyCommand}}{{ end }}
{{ end }}
#### DO NOT ADD ANY ENTRIES BELOW THIS LINE! ####
#### Values below will be added to all above Host entries unless specifically overridden ####
#### Customized/Additional SSH Host entries should go in personal_config file ####

Host *
 ForwardAgent yes
 ForwardX11Trusted yes
 ServerAliveInterval 30
 ServerAliveCountMax 5
 HostKeyAlgorithms +ssh-rsa
 PubkeyAcceptedKeyTypes +ssh-rsa
 UseKeyChain yes
 AddkeysToAgent yes
 User {{ .Username }}
 Port 22

