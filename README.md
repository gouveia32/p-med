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
    <form id='form' name='form' class='form-horizontal dataForm' action='/admin/import/consulta' method='post'	enctype='multipart/form-data'>	<label class='label-micro'>Paciente: {{.paciente.Nome}}</label>	<hr><h2 class="ql-align-center"><strong><u>Consulta</u></strong></h2><p class="ql-align-center"><br></p><p><strong>Anaminese:</strong></p><ol><li>Tem <strong>Febre? </strong> 
        <input type='radio' title='Tem Febre' name='cpo0' id='febre' onchange='atualiza(febre,this.value)' > Sim </input>
        <input type='radio' title='Tem Febre' name='cpo0' id='febre' onchange='atualiza(febre,this.value)' > Não </input>
        </li><li><strong>Diarréia?   </strong> 
        <input type='radio' title='Diarréia' name='cpo1' id='diarreia' onchange='atualiza(diarreia,this.value)' > Sim </input>
        <input type='radio' title='Diarréia' name='cpo1' id='diarreia' onchange='atualiza(diarreia,this.value)' > Não </input>
        </li><li>Já esteve <strong>Internado?</strong>  
        <input type='radio' title='Já esteve Internado' name='cpo2' id='internado' onchange='atualiza(internado,this.value)' > Sim </input>
        <input type='radio' title='Já esteve Internado' name='cpo2' id='internado' onchange='atualiza(internado,this.value)' > Não </input>
        </li><li>Sintoma de <strong>Gripe?</strong> <input type='radio' title='Sintoma gripal' name='cpo3' id='gripe' onchange='atualiza(gripe,this.value)' > Sim </input>
        <input type='radio' title='Sintoma gripal' name='cpo3' id='gripe' onchange='atualiza(gripe,this.value)' > Não </input>
        </li><li>Idade : [[Idade ?idade:text]]</li><li><strong>Fim</strong></li></ol><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p>Aracaju,<input type='text' title='Data' name='data' value='13 de 03 de 2023' ></input>
        </p>
<script src='/static/layui/layui.js?s=36'></script>
<script language='javascript'>
var $,form;

layui.config({
    base : 'js/'
    }).use(['form','element','layer','jquery'],function(){
        form = layui.form;$ = layui.jquery;
    });

function atualiza(key,valor) {
    $('form[name=form]').find('input[name=cpo' + key + ']').val(valor);
    console.log(key);
    console.log(valor);
}

$('.toolTop').click(function () {
    alert('Estou aqui...');
});
</script>