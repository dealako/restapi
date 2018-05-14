package utils

// Art returns the Ascii Art associated with this application
func Art() string {
	return `
                         __                .__
_______   ____   _______/  |______  ______ |__|
\_  __ \_/ __ \ /  ___/\   __\__  \ \____ \|  |
 |  | \/\  ___/ \___ \  |  |  / __ \|  |_> >  |
 |__|    \___  >____  > |__| (____  /   __/|__|
			 \/     \/            \/|__|
`
}

// ArtMessage returns a string with the Ascii Art with a message below
func ArtMessage(message string) string {
	return Art() + "\n\n" + message
}
