<!DOCTYPE html>
<html lang="pt-BR">
<head>
<meta charset="UTF-8">
<title>Redirecionando...</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>
<p>Redirecionando...</p>

<script>
const p = new URLSearchParams(window.location.search);
const destino = p.get("to");

if(destino){
  window.location.replace(destino);
}else{
  document.body.innerHTML = "Link inválido";
}
</script>
</body>
</html>
