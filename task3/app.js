function addSpan(elem) {
    var str = $('<li><span></span> <button class="remove-button">Удалить</button></li>');
    $('span', str).text(elem);
    $('#root ul').append(str);
    $('.remove-button', str).click(function(event){
		$(this).parent().remove()});
}

$(function(){
    $('#root').append('<ul></ul>');
	$('#root').append('<input id="add_task_input">');
	$('#root').append('<button id="add_task">Добавить</button>');
    $('#add_task').click(function(){addSpan($('#add_task_input').val())});
    addSpan('Сделать задание #3 по web-программированию');
})