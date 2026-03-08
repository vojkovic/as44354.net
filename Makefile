LESS_DIR = ./static/less
CSS_DIR = ./static/css

LESS_FILES = $(wildcard $(LESS_DIR)/*.less)
MIN_CSS_FILES = $(patsubst $(LESS_DIR)/%.less, $(CSS_DIR)/%.min.css, $(LESS_FILES))

all: css
	go run .

css: $(MIN_CSS_FILES)

$(CSS_DIR)/%.css: $(LESS_DIR)/%.less
	npx lessc $< $@

$(CSS_DIR)/%.min.css: $(CSS_DIR)/%.css
	npx uglifycss $< > $@
	rm $<

docker:
	docker build -t as44354 .
	docker run -p 8080:8080 as44354

clean:
	rm -f $(CSS_DIR)/*.css $(CSS_DIR)/*.min.css

.PHONY: all css clean docker
