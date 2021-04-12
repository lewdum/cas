## Programação Orientada a Cachaça

Projeto apenas para testes.

As chaves são o checksum de 32 bits do [BLAKE2b](https://en.wikipedia.org/wiki/BLAKE_(hash_function)), mais especificamente o `BLAKE-256`. As chaves são criptograficamente seguras, desde que uma chave secreta seja usada (através do parâmetro `-key`). O servidor não foi testado contra [timing attacks](https://en.wikipedia.org/wiki/Timing_attack). O endpoint `/list` pode ser habilitado pelo parâmetro `-list`. Ele retorna uma lista com todas as chaves no servidor, separada por quebras de linha (`"\n"`). Atualmente, o servidor não usa qualquer técnica de [rate limiting](https://en.wikipedia.org/wiki/Rate_limiting), o que o torna inadequado para uso real.

O cliente contata o servidor sempre no endereço `localhost:8080`. Isso será mudado em breve. O cliente foi projetado em [Svelte](https://svelte.dev/) e usa tecnologias relativamente recentes como [grid layouts](https://www.w3schools.com/css/css_grid.asp), [CSS custom properties](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties), etc. Certifique-se que seu navegador suporta essas tecnologias.
