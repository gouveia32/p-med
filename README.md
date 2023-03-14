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
    <button type='button' class='btn flat btn-info' onclick="tableToExcel('copyTable', 'Player_scores')">
		Processar
	</button>
	<button type='reset' class='btn flat btn-default dataFormReset'>
		Restaurar
	</button>
    <hr><h2 class="ql-align-center"><strong><u>Consulta</u></strong></h2><p class="ql-align-center"><br></p><p><strong>Anaminese:</strong></p><ol>
      <li>Tem <strong>Febre? </strong> 
        <input type='radio' title='Tem Febre' name='cpo0' id='cpo00' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Tem Febre' name='cpo0' id='cpo01' onchange='atualiza(this.name,this.value)' > Não </input>
      </li>
      <li><strong>Diarréia?   </strong> 
        <input type='radio' title='Diarréia' name='cpo1' id='cpo10' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Diarréia' name='cpo1' id='cpo11' onchange='atualiza(this.name,this.value)' > Não </input>
      </li>
      <li>Já esteve <strong>Internado?</strong>  
        <input type='radio' title='Já esteve Internado' name='cpo2' id='cpo20' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Já esteve Internado' name='cpo2' id='cpo21' onchange='atualiza(this.name,this.value)' > Não </input>
      </li>
      <li>Sintoma de <strong>Gripe?</strong> 
        <input type='radio' title='Sintoma gripal' name='cpo3' id='cpo30' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Sintoma gripal' name='cpo3' id='cpo31' onchange='atualiza(this.name,this.value)' > Não </input>
      </li>
      <li> <strong>Idade : </strong>
        <input type='text' title='Idade' name='cpo1'></input>
      </li>
      <li><strong>Fim</strong></li></ol><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p>Aracaju,
        <input type='text' title='Data' name='data' value='13 de 03 de 2023' ></input>
        </p>
<script src='/static/layui/layui.js?s=36'></script>
<script language='javascript'>
var $,form;

layui.config({
    base : 'js/'}).use(['form','element','layer','jquery'],function(){
        form = layui.form;
        $ = layui.jquery;
});

function atualiza(name,value) {
    $('form[name=form]').find('input[name=' + name + ']').val(value);	
    console.log(name);	
    console.log(value);
}

var tableToExcel = (function() {
  
  var uri = 'data:application/vnd.ms-excel;base64,',
    template = '<html xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns="http://www.w3.org/TR/REC-html40"><head><!--[if gte mso 9]><xml><x:ExcelWorkbook><x:ExcelWorksheets><x:ExcelWorksheet><x:Name>{worksheet}<\/x:Name><x:WorksheetOptions><x:DisplayGridlines/><\/x:WorksheetOptions><\/x:ExcelWorksheet><\/x:ExcelWorksheets><\/x:ExcelWorkbook><\/xml><![endif]--><\/head><body><table>{table}<\/table><\/body><\/html>',
    base64 = function(s) {
      return window.btoa(unescape(encodeURIComponent(s)))
    },
    format = function(s, c) {
      return s.replace(/{(\w+)}/g, function(m, p) {
        return c[p];
      })
    }
  return function tableToExcel(table, name) {
    
    if (!table.nodeType) {
    table = document.getElementById(table)}
    var cloned = $('#copyTable').clone().appendTo('.hidden_table')
    cloned.find('input[type="radio"]:not(:checked) + span').remove();
   
    var ctx = {
      worksheet: name || 'Worksheet',
      table: cloned.html()
    }
    cloned.remove();
    //window.location.href = uri + base64(format(template, ctx))
    console.log(format(template, ctx));
  }
})();    

</script>