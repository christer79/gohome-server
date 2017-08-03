Name:		gohome-server
Version:	0.0.3
Release:	1%{?dist}
Summary:	Binary for running home server

License:	GPL

%description
Binary for running home server

%prep

%build
go build

%install
mkdir -p %{buildroot}/opt/gohome-server/
mkdir -p %{buildroot}/etc/gohome-server/
mkdir -p %{buildroot}/usr/lib/systemd/system/

cp -r html %{buildroot}/opt/gohome-server/
cp gohome-server %{buildroot}/opt/gohome-server/
cp config.yml %{buildroot}/etc/gohome-server/
cp gohome-server.service %{buildroot}/usr/lib/systemd/system/

%files
%defattr(-,root,root,-)
/opt/gohome-server/html
%attr(755,root,root) /opt/gohome-server/gohome-server
%attr(755,root,root) /usr/lib/systemd/system/
%config(noreplace) /etc/gohome-server/config.yml

%doc README.adoc



%changelog
