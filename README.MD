# RagulesaSkins

A simple library for converting Minecraft skins: Bedrock Edition, with automatic image size detection. 

Based on https://github.com/Encritary/BedrockSkinConverter 💖

## Usage
 
`ImageToSkinData(image *image.NRGBA)` Converts an image to Bedrock's SkinData (`string`)

`SkinDataToImage(skinData string)` Converts Bedrock's SkinData to an image (`*image.NRGBA`)