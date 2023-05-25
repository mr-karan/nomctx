# Source code for the VHS nomctx demo.
#
# To run:
#
#     vhs < demo.tape

Output demo.gif

Set TypingSpeed 75ms
Set Shell zsh

Set Width 1400
Set Height 700
Set Theme "Catppuccin Mocha"

Type@40ms "echo 'Welcome to nomctx! A tool for switching across multiple Nomad clusters.'"
Sleep 300ms
Enter
Sleep 2s

# Clear screen
Hide
Type "clear"
Enter
Show

Type "cat ~/.nomctx/config.hcl"
Sleep 300ms
Enter
Sleep 2s

# Clear screen
Hide
Type "clear"
Enter
Show

Type "nomctx"
Sleep 500ms
Enter
Sleep 1s
Up 1
Sleep 300ms
Up 1
Sleep 300ms
Down 2
Sleep 500ms
Enter
Sleep 2s

# Clear screen
Hide
Type "clear"
Enter
Show

Type "nomctx current-context"
Sleep 500ms
Enter
Sleep 2s

# Clear screen
Hide
Type "clear"
Enter
Show

Type "nomctx switch-namespace"
Sleep 500ms
Enter
Sleep 1s
Up 4
Sleep 500ms
Down 4
Sleep 1s
Up
Enter
Sleep 2s
