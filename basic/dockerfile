FROM alpine:3.5
LABEL AUTHOR="J.-Bajoul Kakaei"

ENV PORT 5050

COPY ./Udemy_Cloud_Native_Go /app/Udemy_Cloud_Native_Go
RUN chmod +x /app/Udemy_Cloud_Native_Go

EXPOSE ${PORT}
ENTRYPOINT /app/Udemy_Cloud_Native_Go