FROM postgres:17.4

# Set the locale to Japanese
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8

# Set the timezone to Tokyo
ENV TZ Asia/Tokyo