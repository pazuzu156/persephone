@extends('layout.main')

@section('content')
@include('docs.nav')

<div class="col-lg-10 ml-auto cmd-div">
    @php
    $file = resource_path().'/views/docs/markdown/help.md';
    $handle = fopen($file, 'r');
    $content = fread($handle, filesize($file));
    fclose($handle);

    echo Markdown::parse($content)
    @endphp
</div>

@endsection
