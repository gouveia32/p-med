<section class='content'>
    <form id='form' name='form' class='form-horizontal dataForm' action='/admin/import/consulta' method='post'	enctype='multipart/form-data'>	<label class='label-micro'>Paciente: {{.paciente.Nome}}</label>	<hr><h2 class="ql-align-center"><strong><u>Consulta</u></strong></h2><p class="ql-align-center"><br></p><p><strong>Anaminese:</strong></p><ol><li>Tem <strong>Febre? </strong> 
        <input type='radio' title='Tem Febre' name='cpo0' id='cpo00' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Tem Febre' name='cpo0' id='cpo01' onchange='atualiza(this.name,this.value)' > Não </input>
        </li><li><strong>Diarréia?   </strong> 
        <input type='radio' title='Diarréia' name='cpo1' id='cpo10' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Diarréia' name='cpo1' id='cpo11' onchange='atualiza(this.name,this.value)' > Não </input>
        </li><li>Já esteve <strong>Internado?</strong>  
        <input type='radio' title='Já esteve Internado' name='cpo2' id='cpo20' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Já esteve Internado' name='cpo2' id='cpo21' onchange='atualiza(this.name,this.value)' > Não </input>
        </li><li>Sintoma de <strong>Gripe?</strong> 
        <input type='radio' title='Sintoma gripal' name='cpo3' id='cpo30' onchange='atualiza(this.name,this.value)' > Sim </input>
        <input type='radio' title='Sintoma gripal' name='cpo3' id='cpo31' onchange='atualiza(this.name,this.value)' > Não </input>
        </li><li>Idade : [[Idade ?idade:text]]</li><li><strong>Fim</strong></li></ol><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p>Aracaju,
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
    $('form[name=form]').find('input[name=' + name + ']').val(1);	
        console.log(name);	
        console.log(value);
    }
    
    $('.toolTop').click(function () {
        alert('Estou aqui...');
    });
</script>