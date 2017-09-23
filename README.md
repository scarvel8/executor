# executor
Execute remote commands onto remote clients (golang)

Usage:

executor exec (command)
Execute the command locally on the server in exec style

executor exec --shell (shell) (command)
Execute the command with a specific previously defined shell

Next steps:

1.) Define different shells (sh, bash, cmd, powershell) for flags
2.) Parse JSON input for cmd and shell
3.) Web service
4.) Expose HTTP API and accept JSON POST for command running
5.) Define config files (HCL?)
6.) Move shell definition to a config file
7.) Begin remote executor work, define types for remote server definition, etc.