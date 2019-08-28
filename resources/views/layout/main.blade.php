<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Persephone</title>
        <link rel="stylesheet" href="{{ asset('css/app.css') }}">
    </head>
    <body>
        <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
            <div class="container">
                <a href="#" class="navbar-brand">Persephone</a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarDropMenu" aria-controls="navbarDropMenu" aria-expanded="false" aria-label="Toggle Navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarDropMenu">
                    <div class="navbar-nav ml-auto">
                        <a href="{{ url('/') }}" class="nav-link nav-item active">Home</a>
                        <a href="{{ route('auth.login') }}" class="nav-link nav-item">Login</a>
                    </div>
                </div>
            </div>
        </nav>
        <div class="container mt-md-3">
            <div class="content">
                @if(isset($pageTitle))
                <h1>{{ $pageTitle }}</h1>
                @endif

                @yield('content')
            </div>
        </div>
        <nav class="navbar navbar-expanded-lg navbar-dark bg-dark fixed-bottom">
            <div class="container navbar-text">
                <div class="col-lg-10">
                    Made for Untrodden Corridors of Hades.
                    <a href="https://discord.gg/{{ env('DISCORD_GUILD_INVITE_CODE') }}"><i class="fab fa-discord"></i></a> |
                    <a href="https://github.com/pazuzu156/persephone"><i class="fab fa-github"></i></a>
                </div>
                <div class="col-lg-2 text-right">
                    &copy; 2019 <a href="https://kalebklein.com">Kaleb Klein</a>.
                </div>
            </div>
        </nav>
        <script src="{{ asset('js/app.js') }}"></script>
        @if(Session::has('alert'))
            <script>$.notify("{{ Session::get('message') }}", {
                type: "{{ Session::get('alert') }}"
            })</script>
        @endif
    </body>
</html>
