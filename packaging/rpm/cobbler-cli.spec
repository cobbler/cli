#
# spec file for package cobbler-cli
#
# Copyright (c) 2024 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           cobbler-cli
Version:        0.0.1
Release:        0
Summary:        A standalone CLI for the Cobbler Daemon
License:        GPL-2.0-or-later
URL:            https://github.com/cobbler/cli
Source0:        %{name}-%{version}.tar.gz
Source1:        vendor.tar.gz
BuildRequires:  golang(API) >= 1.22
BuildRequires:  zsh
BuildRequires:  fish
BuildRequires:  bash-completion

%description
Independent CLI written in Go for the Cobbler server.

%package bash-completion
Summary:        Bash Completion for %{name}
Group:          System/Shells
Requires:       %{name} = %{version}
Requires:       bash-completion
Supplements:    (%{name} and bash-completion)
BuildArch:      noarch
Provides:       %{name}-bash-completion = %{version}
Obsoletes:      %{name}-bash-completion < %{version}
Conflicts:      %{name}-bash-completion

%description bash-completion
Bash command line completion support for %{name}.

%package zsh-completion
Summary:        Zsh Completion for %{name}
Group:          System/Shells
Requires:       %{name} = %{version}
Requires:       zsh
Supplements:    (%{name} and zsh)
BuildArch:      noarch
Provides:       %{name}-zsh-completion = %{version}
Obsoletes:      %{name}-zsh-completion < %{version}
Conflicts:      %{name}-zsh-completion

%description zsh-completion
Zsh command line completion support for %{name}.

%package fish-completion
Summary:        Fish completion for %{name}
Group:          System/Shells
Requires:       %{name} = %{version}
Requires:       fish
Supplements:    (%{name} and fish)
BuildArch:      noarch
Provides:       %{name}-fish-completion = %{version}
Obsoletes:      %{name}-fish-completion < %{version}
Conflicts:      %{name}-fish-completion

%description fish-completion
Fish command line completion support for %{name}.

%prep
%autosetup -p1
%autosetup -T -D -a 1

%build
go build \
   -mod=vendor \
   -buildmode=pie \
   -o cobbler
make shell_completions

%install
install -D -m0755 cobbler %{buildroot}%{_bindir}/cobbler
# Shell completions
install -D -m0644 config/completions/bash/cobbler "%{buildroot}%{_datarootdir}/bash-completion/completions/cobbler"
install -D -m0644 config/completions/zsh/cobbler "%{buildroot}%{_sysconfdir}/zsh_completion.d/_cobbler"
install -D -m0644 config/completions/fish/cobbler "%{buildroot}/%{_datadir}/fish/vendor_completions.d/cobbler.fish"

%files
%license LICENSE
%doc README.md
%{_bindir}/cobbler

%files bash-completion
%defattr(-,root,root)
%{_datarootdir}/bash-completion/completions/cobbler

%files zsh-completion
%defattr(-,root,root)
%{_sysconfdir}/zsh_completion.d/_cobbler

%files fish-completion
%defattr(-,root,root)
%{_datadir}/fish/vendor_completions.d/cobbler.fish

%changelog
