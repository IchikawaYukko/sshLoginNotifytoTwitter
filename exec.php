<?php
  echo exec("last -ain 1|sed 's/.*in[^0-9]*\([0-9.]*\)[^0-9]*.*/\\1/'|head -n 1");
?>
