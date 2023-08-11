# sistema de Prontuário Médico p-med

# Modelo: formato do campo = [[Prompt?name:type=Answer]]
  Exemplos: [[Sexo?sexo:escolha=M,F]]
            [[Data atual:hoje=data]] ou [[Data atual:hoje=data_extenso]]
            [[Exame?exame:texto_auto=resp1,resp2,resp3]]
            [[Exame?exame:texto_auto=.um,dois,tres]]  busca lista um,dois,tres
            [[Exame 2?exame2:lista=exame]]  busca lista de autopreenchimento da tabela lista

            [[Exame?exame1:texto_auto=RX da Face, Sumário de Urina]]
            [[Exame?exame2:texto_auto=TC,Ressonância]]





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
    <hr><h2 class="ql-align-center"><strong><u>Consulta</u></strong></h2><p class="ql-align-center"><br></p><p><strong>Anaminese:</strong></p><ol><li>Tem <strong>Febre? </strong> 
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
        </li><li> <strong>Idade : </strong>
        <input type='text' title='Idade' name='cpo1'></input></li><li><strong>Fim</strong></li></ol><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p><br></p><p>Aracaju,
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

























Imports System.Collections.Generic
Imports MySql.Data.MySqlClient

Namespace DAL
    Public Class ClassBD
        Private conn As MySqlConnection
        Private comando As MySqlCommand
        Private comandoTemp As MySqlCommand
        Private adap As MySqlDataAdapter
        Private strConn As String

        Public Shared servidor As [String] = "localhost"
        Public Shared banco As [String] = "ccb"
        Public Shared usuario As [String] = "root"
        Public Shared senha As [String] = "123456"
        Public Shared SessaoTimeOut As [Int32] = 100
        Public Shared StringDeConexao As [String]

        Public Sub New()
            Me.strConn = StringDeConexao
            conn = New MySqlConnection(Me.strConn)
        End Sub

        Public Shared Function create() As ClassBD
            Return New ClassBD()
        End Function

        ' code in vb.net you can convert it to c# easily
Dim ConnectionString As String = "your connection string"
Public Function KillSleepingConnections(Optional p1 As Integer = 0) As Integer
    Dim query As String = "SHOW FULL PROCESSLIST"
Dim mySqlDataReader As MySqlDataReader
    '
    Try
        'ConnectDatabase()
        Dim mySqlConnection As MySqlConnection = New MySqlConnection(strConn)
        ' 2. open connection to database using connectionString
        mySqlConnection.Open()
        ' 3. Create a command to hold values
        Dim objCmd As New MySqlCommand(query, mySqlConnection)
        ' 4. Add parameters for sqlCommand
        mySqlDataReader = objCmd.ExecuteReader()
        If mySqlDataReader.HasRows Then
            Do While mySqlDataReader.Read()
                ' kill processes with elapsed time > 200 seconds and in Sleep 
                If mySqlDataReader.GetInt32(5) > 200 And mySqlDataReader.GetString(4) = "Sleep" Then
                    KillMySqlProcess("KILL " & mySqlDataReader.GetInt32(0))
                End If
            Loop
        End If
        If Not mySqlDataReader Is Nothing Then mySqlDataReader.Close()
        If Not mySqlConnection Is Nothing Then mySqlConnection.Close()
    Catch ex As MySqlException
        Return -1
    End Try
    Return 0
End Function

Public Function KillMySqlProcess(ByVal myQuery As String) As Integer
    '1. Create a query
    Dim query As String = myQuery
    '
    Try
        Dim mySqlConnection As MySqlConnection = New MySqlConnection(ConnectionString)
        ' 2. open connection to database using connectionString
        mySqlConnection.Open()
        ' 3. Create a command to hold values
        Dim objCmd As New MySqlCommand(query, mySqlConnection)
        objCmd.ExecuteNonQuery()
        mySqlConnection.Close()
    Catch ex As MySqlException
        Return -1
    End Try

    Return 0
End Function




       Public Function KillSleepingConnections1(ByVal iMinSecondsToExpire As Integer) As Integer
            Dim strSQL As String = "show processlist"
            Dim m_ProcessesToKill As System.Collections.ArrayList = New ArrayList()
            'Dim myConn As MySqlConnection
            Dim myCmd As MySqlCommand
            Dim MyReader As MySqlDataReader = Nothing

            Try
                conn = New MySqlConnection(Me.strConn)
                myCmd = New MySqlCommand(strSQL, conn)
                conn.Open()

                MyReader = myCmd.ExecuteReader()

                While MyReader.Read()
                    Dim iPID As Integer = Convert.ToInt32(MyReader("Id").ToString())
                    Dim strState As String = MyReader("Command").ToString()
                    Dim iTime As Integer = Convert.ToInt32(MyReader("Time").ToString())

                    'If strState = "Sleep" AndAlso iTime >= iMinSecondsToExpire AndAlso iPID > 0 Then
                    If strState = "Sleep" AndAlso iPID > 0 Then
                        m_ProcessesToKill.Add(iPID)
                    End If
                End While

                MyReader.Close()

                For Each aPID As Integer In m_ProcessesToKill
                    strSQL = "kill " & aPID
                    myCmd.CommandText = strSQL
                    myCmd.ExecuteNonQuery()
                Next

            Catch excep As Exception
            Finally

                If MyReader IsNot Nothing AndAlso Not MyReader.IsClosed Then
                    MyReader.Close()
                End If

                If conn IsNot Nothing AndAlso conn.State = ConnectionState.Open Then
                    conn.Close()
                End If
            End Try

            Return m_ProcessesToKill.Count
        End Function

        Public Function exePesquisa(query As String, Optional Param As List(Of MySqlParametro) = Nothing) As DataTable
            Dim dados As New DataTable()

            'executa a query
            Try
                Me.adap = New MySqlDataAdapter(query, StringDeConexao)

                'insere os parametros na query
                If Param IsNot Nothing Then
                    For Each p As MySqlParametro In Param
                        Me.adap.SelectCommand.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If

                Me.adap.Fill(dados)
            Catch erro As Exception
                MessageBox.Show("ERRO: " & erro.ToString)
            Finally
                'Me.adap.Dispose()
                 KillSleepingConnections(SessaoTimeOut)
                'DisconnectDatabase()
            End Try

            Return dados
        End Function

        Public Function exeNonQuery(query As String, Optional param As List(Of MySqlParametro) = Nothing) As Integer
            Dim id As Integer = 0
            Try
                'KillSleepingConnections(100)
                ConnectDatabase()

                Me.comando = New MySqlCommand(query, Me.conn)
                'adiciona parametros
                If param IsNot Nothing Then
                    For Each p As MySqlParametro In param
                        Me.comando.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If
                'executa a query
                Me.comando.ExecuteNonQuery()
                'Retorna o novo Id
                id = Convert.ToInt32(comando.LastInsertedId)
            Catch erro As Exception
                Throw New Exception("ERRO: " + erro.Message)
            Finally
                'Me.comando.Dispose()
                'DisconnectDatabase()
                 KillSleepingConnections(SessaoTimeOut)
            End Try

            Return id
        End Function

        Public Function exeScalar(query As String, Optional param As List(Of MySqlParametro) = Nothing) As Integer
            Dim resultado As Integer = 0

            Try
                'KillSleepingConnections(100)
                ConnectDatabase()

                Me.comando = New MySqlCommand(query, conn)
                'adiciona parametros
                If param IsNot Nothing Then
                    For Each p As MySqlParametro In param
                        Me.comando.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If
                'executa a query
                Dim result As Object = comando.ExecuteScalar()
                If result IsNot Nothing Then
                    resultado = Convert.ToInt32(result)
                End If
            Catch erro As Exception
                Throw New Exception("ERRO: " + erro.Message)
            Finally
                'Me.comando.Dispose()
                'DisconnectDatabase()
                 KillSleepingConnections(SessaoTimeOut)
            End Try
            Return resultado
        End Function

        <CLSCompliant(False)> Public Function Reader(query As String, Optional param As List(Of MySqlParametro) = Nothing) As MySqlDataReader
            Dim registro As MySqlDataReader = Nothing

            Try
                'Me.conn = New MySqlConnection(strConn)
                'Me.conn.Open()
                ConnectDatabase()

                Me.comandoTemp = New MySqlCommand(query, Me.conn)

                If param IsNot Nothing Then
                    For Each p As MySqlParametro In param
                        comandoTemp.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If
                registro = comandoTemp.ExecuteReader()
            Catch erro As Exception
                Throw New Exception("ERRO: " + erro.Message)
             Finally
                'DisconnectDatabase()
                 KillSleepingConnections(SessaoTimeOut)
            End Try
            Return registro
        End Function

        <CLSCompliant(False)> Public Function Fill(query As String, Optional Param As List(Of MySqlParametro) = Nothing) As MySqlDataAdapter
            Dim dados As New DataTable()

            'executa a query
            Try
                ConnectDatabase()
                Me.adap = New MySqlDataAdapter(query, Me.strConn)

                'insere os parametros na query
                If Param IsNot Nothing Then
                    For Each p As MySqlParametro In Param
                        Me.adap.SelectCommand.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If
            Catch erro As Exception
                Throw New Exception("ERRO: " & erro.ToString)
            Finally
                adap.Dispose()
                'DisconnectDatabase()
            End Try

            Return adap
        End Function

        Public Function TestarConexao() As Boolean
            Try

                ConnectDatabase()
                Return True
            Catch erro As Exception
                Throw New Exception("Falha na conexão com o BD!")
            End Try
            'Me.comando.Dispose()
            'DisconnectDatabase()

        End Function

        <CLSCompliant(False)> Public Sub FecharReader(ByVal r As MySqlDataReader)
            If r Is Nothing Then Exit Sub
            DisconnectDatabase()
            Me.comandoTemp.Dispose()
            'conn.Dispose()
        End Sub

        Public Sub ConnectDatabase()
            Try
                If Me.conn.State = ConnectionState.Open Then
                    return
                End If
                'Me.conn.ConnectionString = Me.strConn
                Me.conn.Open()
                'MsgBox("Banco conectado.", MsgBoxStyle.Exclamation)
            Catch erro As Exception
                Throw New Exception("ERRO: " & erro.ToString)
                End
            End Try
        End Sub

        Public Sub DisconnectDatabase()
            Try
                conn.Close()
                conn.Dispose
                conn.ClearAllPoolsAsync()
                KillSleepingConnections(SessaoTimeOut)
                'MsgBox("Banco DESCONECTADO.", MsgBoxStyle.Exclamation)
            Catch myerror As MySql.Data.MySqlClient.MySqlException
            End Try
        End Sub

        Public Function getTabelaViaAdapter(consultaSQL As String, Optional Param As List(Of MySqlParametro) = Nothing) As DataTable
            Try
                ConnectDatabase()
                Dim dt As New DataTable()
                Dim da As MySqlDataAdapter = New MySqlDataAdapter(consultaSQL, conn)

                'insere os parametros na query
                If Param IsNot Nothing Then
                    For Each p As MySqlParametro In Param
                        da.SelectCommand.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If

                da.Fill(dt)

                getTabelaViaAdapter = dt

            Catch erro As Exception
                Throw New Exception("ERRO: " + erro.Message)
            Finally
                'DisconnectDatabase()
            End Try
        End Function

        Public Function getTabelaViaReader(ByVal consultaSQL As String, Optional Param As List(Of MySqlParametro) = Nothing) As DataTable
            Try
                ConnectDatabase()
                Dim dt As New DataTable()
                Dim cmd As New MySqlCommand(consultaSQL, conn)

                'insere os parametros na query
                If Param IsNot Nothing Then
                    For Each p As MySqlParametro In Param
                        cmd.Parameters.AddWithValue(p.Parametro, p.Valor)
                    Next
                End If

                Dim reader As MySqlDataReader = cmd.ExecuteReader()

                dt.Load(reader)

                getTabelaViaReader = dt

            Catch ex As Exception
                Throw ex
            Finally
                'DisconnectDatabase()
                 KillSleepingConnections(SessaoTimeOut)
            End Try
        End Function

        Public Function getDataSet(ByVal consultaSQL As String) As DataSet
            Try
                ConnectDatabase()
                Dim dadapter As New MySqlDataAdapter(consultaSQL, conn)
                Dim dset As New DataSet()
                dadapter.Fill(dset)
                Return dset
            Catch ex As Exception
                Throw ex
            Finally
                'conn.Close()
                'conn.Dispose
                'conn.ClearAllPoolsAsync()
            End Try
        End Function

    End Class

End Namespace







produção
go build -ldflags "-s -w"