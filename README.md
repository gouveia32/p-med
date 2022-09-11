# pm   forked by CCB_go
Uma estrutura de gerenciamento de conteúdo baseada em beego + layui + mysql.

    Observe a descrição
    Eu acabei de aprender a linguagem go por um curto período de tempo. Em função do aprendizado e da prática, este framework possui algumas funções básicas.
    Principalmente funções básicas como login, lista, envio de formulários, etc. Mais tarde, irei melhorar gradualmente o conteúdo ao qual desejo corresponder.
    Se você tiver alguma dúvida, pode fazer perguntas

# Introdução
CCB é um sistema de gerenciamento leve baseado no Beego que é fácil de usar, expandir e possui uma interface amigável.
A estrutura de front-end é baseada em layui para integração de recursos.
Este sistema é desenvolvido com base no beego, usa o banco de dados mysql por padrão e caches redis
# Use tecnologia
Linguagem de back-end: golang
Estrutura de back-end: [beego.me] (http://beego.me).
Estrutura do plano de fundo: [modelo de página única de plano de fundo do nepadmin] (https://github.com/fanjyy/nepadmin).
# Interface de exibição
! [Página inicial] (https://github.com/fengke549015/pm/blob/master/static/images/demo/homepage.jpg)
! [Login] (https://github.com/fengke549015/pm/blob/master/static/images/demo/login.png)
! [Configuração do site] (https://github.com/fengke549015/pm/blob/master/static/images/demo/Website Configuration.png)
! [Envio de formulário] (https://github.com/fengke549015/pm/blob/master/static/images/demo/form submit.png)
# método de instalação    
1. Acesse github.com/fengke549015/pm
2. Crie um banco de dados mysql e importe ccbBeego.sql
3. Modifique data / conf / database.conf para configurar o banco de dados
4. Execute go build
5. Execute ./run.sh start | stop