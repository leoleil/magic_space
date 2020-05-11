jQuery Templates plugin 1.0.1
====
This is the fork of jQuery Templates that we use internally here at Kanban Solutions. We have fixed all known and encountered bugs as well as changing the syntax from the standard

Docs
====

[.tmpl()](http://api.jquery.com/tmpl/)
[jQuery.tmpl()](http://api.jquery.com/jquery.tmpl/)
[.template()](http://api.jquery.com/template/)
[jQuery.template()](http://api.jquery.com/jQuery.template/)

{%= fieldNameOrExpression %}
----
Used for insertion of data values in the rendered template. Evaluates the specified field (property) on the current data item, or the specified JavaScript function or expression.

_identical to the `${}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-equal/_

{%html fieldNameOrExpression %}
----
Used for insertion of HTML markup strings in the rendered template. Evaluates the specified field on the current data item, or the specified JavaScript function or expression.

_identical to the `{{html}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-html/_

{%if fieldNameOrExpression %}
----
Used for conditional insertion of content. Renders the content between the opening and closing template tags only if the specified data item field, JavaScript function or expression does not evaluate to false (or to zero, null, type "undefined", or the empty string ).

_identical to the `{{if}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-if/_

{%elif fieldNameOrExpression %}
----
Used in association with the `{{if}}...{{/if}}` tag to provide alternative content based on the values of one or more expressions. The `{%elif%}` tag can be used with a parameter, as in: `{%if a%}...{%elif b%}...{%/if%}`.

_Added because the old system actualy overloaded the else tag to handle elif as well and I didnt like that_

{%else%}
----
Used in association with the `{{if}}...{{/if}}` tag to provide alternative content based on the values of one or more expressions. The `{%else%}` tag can be used without a parameter, as in: `{%if a%}...{%elif b%}...{%else%}...{%/if%}`

_almost identical to the `{{else}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-else/ but I removed the ability to do `{{else _test_}}`_


{%each( index, value ) collection %}
----
Used to iterate over a data array, and render the content between the opening and closing template tags once for each data item.

_identical to the `{{each}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-each/_


{%tmpl( [data], [options] ) template %}
----
Used to iterate over a data array, and render the content between the opening and closing template tags once for each data item.

_identical to the `{{tmpl}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-tmpl/_


{%wrap( [data], [options] ) template %}
----
Used for composition of templates which incorporate wrapped HTML content. Rendered template items can combine wrapped HTML content with template markup.

_identical to the `{{wrap}}` tag of the origional jQuery Templates http://api.jquery.com/template-tag-wrap/_
