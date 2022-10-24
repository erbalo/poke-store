create or replace function public.uuid_generate_v4()
    returns uuid
    language c
    parallel safe strict
as '$libdir/uuid-ossp', $function$uuid_generate_v4$function$;

create table if not exists products (
    id bigserial primary key not null,
    sku uuid default uuid_generate_v4(),
    "name" varchar(150) not null,
    price numeric(15,2) default 0,
    "status" varchar(20) default "UNKNOWN",
    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone default (now() at time zone 'utc')
)

create index on products(sku);
create index on products(created_at);

create table if not exists inventories (
    id bigserial primary key not null,
    product_id uuid references products(id) not null,
    fulfillment_date date not null,
    quantity bigint default 0, 
    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone default (now() at time zone 'utc'),
    unique(product_id, fulfillment_date)
);

create index on inventories(product_id);
create index on inventories(fulfillment_date);
create index on inventories(created_at);