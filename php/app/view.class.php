<?php
class View {

    private $template_content;

    public function __construct($template_path) {
        $this->template_content = file_get_contents($template_path);
    }

    public function apply($flag, $value) {
        $this->template_content = str_replace('{'.$flag.'}', $value, $this->template_content);
    }

    public function echo() {
        echo $this->template_content;
    }
}

?>