<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=header"/>

![hunteheader](/img/line.png)

<p align="center">
	<a href="https://www.python.org/"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/line/releases"><img src="https://img.shields.io/github/release/joaoviictorti/line?color=blue"></a>
</p>

<h4 align="center">line realizar parse de informações duplicadas</h4>


<p align="center">
  <a href="#características">Características</a> •
  <a href="#instalação">Instalação</a> •
  <a href="#forma-de-utilização"> Forma de utilização</a> •
  <a href="#detalhes">Detalhes</a> •
  <a href="#executando-line">Executando line</a>  
</p>

---


O line é uma ferramenta que tem o objetivo de realizar o parse de várias linhas. É uma excelente opção para ser utilizada em testes de penetração no reconhecimento de aplicações e também em atividades do dia a dia.

Projetei o `line` e mantive um modelo consistentemente passivo para torná-lo útil para pentest e profissionais de TI.

# Características

 - Realizar parse de linhas duplicadas
 - Remover espaços entre as linhas
 - Remover linhas vazias
 - Ignora case-sensitive

# Forma de utilização

- Enviando o output para um arquivo
```sh
cat urls.txt | line -f parse_urls.txt
```

- Ignorando case-sensitive
```sh
cat urls.txt | line -if parse_urls.txt
```

- Removendo linhas vazias
```sh
cat urls.txt | line -ef parse_urls.txt
```

- Pode usar todas as flags agrupadas
```sh
cat urls.txt | line -stif parse_urls.txt
```

# Detalhes

![line](/img/help_line.png)


# Instalação

line requer o **golang** instalado, para usar:

```sh
go install -v github.com/joaoviictorti/line/cmd/line@latest
```

# Executando line

![line](/img/exec.png)


<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=footer"/>