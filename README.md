# sistema de Prontuário Médico p-med




ALTER DATABASE `p-med` 
  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci; 
  
ALTER TABLE nota 
  CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; 
  
-- Change a column
ALTER TABLE nota
  CHANGE conteudo conteudo TEXT
  CHARACTER SET utf8mb4 COLLATE UTF8MB4_UNICODE_CI;
  
ALTER TABLE atendimento 
  CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; 
  
-- Change a column
ALTER TABLE atendimento
  CHANGE conteudo conteudo TEXT
  CHARACTER SET utf8mb4 COLLATE UTF8MB4_UNICODE_CI;  









<section class='content'>
    <form id='form' name='form' class='form-horizontal dataForm' action='/admin/import/consulta' method='post'	enctype='multipart/form-data'>	
        <label class='label-micro'>Paciente: {{.paciente.Nome}}</label>	
            <hr><h2 class="ql-align-center"><strong><u>Consulta</u></strong></h2><p class="ql-align-center"><br></p><p><strong>Anaminese:</strong></p><ol><li>Tem <strong>Febre? </strong> 
            <input type='radio' title='Tem Febre' name='cpo0' id='febre' onchange='atualiza(this.id,this,this.value)' > Sim </input>
            <input type='radio' title='Tem Febre' name='cpo0' id='febre' onchange='atualiza(this.id,this,this.value)' > Não </input>
            </li><li><strong>Diarréia?   </strong> 
            <input type='radio' title='Diarréia' name='cpo1' id='diarreia' onchange='atualiza(this.id,this,this.value)' > Sim </input>
            <input type='radio' title='Diarréia' name='cpo1' id='diarreia' onchange='atualiza(this.id,this,this.value)' > Não </input>
            </li><li>Já esteve <strong>Internado?</strong>  
            <input type='radio' title='Já esteve Internado' name='cpo2' id='internado' onchange='atualiza(this.id,this,this.value)' > Sim </input>
            <input type='radio' title='Já esteve Internado' name='cpo2' id='internado' onchange='atualiza(this.id,this,this.value)' > Não </input>
            </li><li><strong>Idade</strong> [[Idade?idade:text]]</li><li><strong>Fim</strong><span class="ql-cursor">﻿</span></li></ol><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p>
            <p>Aracaju,<input type='text' title='Data' name='data' value='13 de 03 de 2023' ></input></p>	
    </form>
</section>

<script src="/static/layui/layui.js?s=36"></script>
<script language="javascript">
    var $,form;
    
    layui.config({
            base : "js/"
        }).use(['form','element','layer','jquery'],function(){
          form = layui.form; //Somente após esta etapa ser realizada, alguns elementos do formulário serão decorados automaticamente com sucesso
          $ = layui.jquery;
          
        });

    function esconder(key,esconder){
        var label = document.getElementById("label" + key); ;
        label.hidden = esconder;   

        var input = document.getElementById("campo" + key); ;
        input.hidden = esconder;   

    }

    function substituir(campo,valor,key) {

        var p1 = $("input[name^=inicio" + key + "]").val();
        var p2 = $("input[name^=fim" + key + "]").val();


        var c = '[['+campo.Descricao+'?'+campo.Nome+':'+campo.Resposta+']]';
        console.log('c:',c);
        bdhtml=quill.root.innerHTML;// Get the... Of the current page html Code
        if (p1 == 0) {
            p1 = bdhtml.indexOf(c); // localiza inicio do campo
            if (p2 == 0) {
                p2 = p1 + c.length;
            }
        }
        console.log(valor);
        console.log(p1);
     
        if (valor != '') {
            if (p1 > 0) {
                $("form[name=form]").find("input[name=inicio" + key + "]").val(p1);
                //console.log(campo.Fim);
            }
            console.log(p1);
            console.log(p2);
            parte1=bdhtml.substring(0,p1); // pega a parte inicial 
            parte2=bdhtml.substring(p2,bdhtml.length); // pega a parte final
            p2 = (+p1) + valor.length;

            //console.log(valor.length);
            //console.log(p2);
            $("form[name=form]").find("input[name=fim" + key + "]").val(p2);

            quill.root.innerHTML= parte1 + valor + parte2;
       }    
    }


</script>
