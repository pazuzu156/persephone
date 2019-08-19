all: persephone

persephone: clean
	NOTAR=true ./build

install:
	sudo install -D -m644 persephone /usr/bin/persephone
	sudo chmod +x /usr/bin/persephone
	mkdir -pv ${HOME}/persephone/temp
	mkdir -pv ${HOME}/persephone/static/fonts
	mkdir -pv ${HOME}/persephone/static/images
	install -D -m644 config.example.json ${HOME}/persephone/config.json
	install -D -m644 static/fonts/NotoSans-Bold.ttf ${HOME}/persephone/static/fonts/NotoSans-Bold.ttf
	install -D -m644 static/fonts/NotoSans-Regular.ttf ${HOME}/persephone/static/fonts/NotoSans-Regular.ttf
	install -D -m644 static/images/background.png ${HOME}/persephone/static/images/background.png

uninstall:
	sudo rm -rf /usr/bin/persephone
	rm -rf ${HOME}/persephone/temp
	rm -rf ${HOME}/persephone/static

clean:
	if [ -f persephone ]; then rm persephone; fi
