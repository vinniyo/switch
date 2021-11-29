# Nintendo Switch OLED stock checker
Program that checks if the switch OLED is available on Nintendo using ChromeDP. 

It will check the pages for the Add Cart button in headless mode (No chrome window) then if the button is found it will open a broswer window, add it to the cart and you checkout from there. 

ChromeDP will fill up your temp directory so every 100 iterations it will clear out the directory on files that are not locked/in-use. Check the source to see the pages it checks. 
