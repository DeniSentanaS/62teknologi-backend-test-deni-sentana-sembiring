PGDMP                         {            enamdua    15.2    15.2 /    3           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            4           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            5           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            6           1262    34100    enamdua    DATABASE     ~   CREATE DATABASE enamdua WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE enamdua;
                postgres    false            �            1259    34167    businesses_categories    TABLE     �   CREATE TABLE public.businesses_categories (
    businesses_entity_id bigint NOT NULL,
    categories_entity_id bigint NOT NULL
);
 )   DROP TABLE public.businesses_categories;
       public         heap    postgres    false            �            1259    34102    businesses_entities    TABLE     ?  CREATE TABLE public.businesses_entities (
    id bigint NOT NULL,
    alias text,
    name text,
    image_url text,
    is_closed boolean,
    url text,
    review_count bigint,
    rating numeric,
    phone text,
    distance numeric,
    create_at timestamp with time zone,
    update_at timestamp with time zone
);
 '   DROP TABLE public.businesses_entities;
       public         heap    postgres    false            �            1259    34101    businesses_entities_id_seq    SEQUENCE     �   CREATE SEQUENCE public.businesses_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.businesses_entities_id_seq;
       public          postgres    false    215            7           0    0    businesses_entities_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.businesses_entities_id_seq OWNED BY public.businesses_entities.id;
          public          postgres    false    214            �            1259    34143    businesses_locations    TABLE     �   CREATE TABLE public.businesses_locations (
    businesses_entity_id bigint NOT NULL,
    locations_entity_id bigint NOT NULL
);
 (   DROP TABLE public.businesses_locations;
       public         heap    postgres    false            �            1259    34119    businesses_transactions    TABLE     �   CREATE TABLE public.businesses_transactions (
    businesses_entity_id bigint NOT NULL,
    transactions_entity_id bigint NOT NULL
);
 +   DROP TABLE public.businesses_transactions;
       public         heap    postgres    false            �            1259    34159    categories_entities    TABLE     �   CREATE TABLE public.categories_entities (
    id bigint NOT NULL,
    alias text NOT NULL,
    title text NOT NULL,
    create_at timestamp with time zone,
    update_at timestamp with time zone
);
 '   DROP TABLE public.categories_entities;
       public         heap    postgres    false            �            1259    34158    categories_entities_id_seq    SEQUENCE     �   CREATE SEQUENCE public.categories_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.categories_entities_id_seq;
       public          postgres    false    223            8           0    0    categories_entities_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.categories_entities_id_seq OWNED BY public.categories_entities.id;
          public          postgres    false    222            �            1259    34135    locations_entities    TABLE     |  CREATE TABLE public.locations_entities (
    id bigint NOT NULL,
    address1 text NOT NULL,
    address2 text,
    address3 text,
    city text NOT NULL,
    zip_code text NOT NULL,
    country text NOT NULL,
    state text NOT NULL,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    create_at timestamp with time zone,
    update_at timestamp with time zone
);
 &   DROP TABLE public.locations_entities;
       public         heap    postgres    false            �            1259    34134    locations_entities_id_seq    SEQUENCE     �   CREATE SEQUENCE public.locations_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 0   DROP SEQUENCE public.locations_entities_id_seq;
       public          postgres    false    220            9           0    0    locations_entities_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE public.locations_entities_id_seq OWNED BY public.locations_entities.id;
          public          postgres    false    219            �            1259    34111    transactions_entities    TABLE     �   CREATE TABLE public.transactions_entities (
    id bigint NOT NULL,
    transaction text NOT NULL,
    create_at timestamp with time zone,
    update_at timestamp with time zone
);
 )   DROP TABLE public.transactions_entities;
       public         heap    postgres    false            �            1259    34110    transactions_entities_id_seq    SEQUENCE     �   CREATE SEQUENCE public.transactions_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 3   DROP SEQUENCE public.transactions_entities_id_seq;
       public          postgres    false    217            :           0    0    transactions_entities_id_seq    SEQUENCE OWNED BY     ]   ALTER SEQUENCE public.transactions_entities_id_seq OWNED BY public.transactions_entities.id;
          public          postgres    false    216            �           2604    34105    businesses_entities id    DEFAULT     �   ALTER TABLE ONLY public.businesses_entities ALTER COLUMN id SET DEFAULT nextval('public.businesses_entities_id_seq'::regclass);
 E   ALTER TABLE public.businesses_entities ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214    215            �           2604    34162    categories_entities id    DEFAULT     �   ALTER TABLE ONLY public.categories_entities ALTER COLUMN id SET DEFAULT nextval('public.categories_entities_id_seq'::regclass);
 E   ALTER TABLE public.categories_entities ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    223    222    223            �           2604    34138    locations_entities id    DEFAULT     ~   ALTER TABLE ONLY public.locations_entities ALTER COLUMN id SET DEFAULT nextval('public.locations_entities_id_seq'::regclass);
 D   ALTER TABLE public.locations_entities ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219    220            �           2604    34114    transactions_entities id    DEFAULT     �   ALTER TABLE ONLY public.transactions_entities ALTER COLUMN id SET DEFAULT nextval('public.transactions_entities_id_seq'::regclass);
 G   ALTER TABLE public.transactions_entities ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    217    217            0          0    34167    businesses_categories 
   TABLE DATA           [   COPY public.businesses_categories (businesses_entity_id, categories_entity_id) FROM stdin;
    public          postgres    false    224   �>       '          0    34102    businesses_entities 
   TABLE DATA           �   COPY public.businesses_entities (id, alias, name, image_url, is_closed, url, review_count, rating, phone, distance, create_at, update_at) FROM stdin;
    public          postgres    false    215   �>       -          0    34143    businesses_locations 
   TABLE DATA           Y   COPY public.businesses_locations (businesses_entity_id, locations_entity_id) FROM stdin;
    public          postgres    false    221   �A       *          0    34119    businesses_transactions 
   TABLE DATA           _   COPY public.businesses_transactions (businesses_entity_id, transactions_entity_id) FROM stdin;
    public          postgres    false    218   �A       /          0    34159    categories_entities 
   TABLE DATA           U   COPY public.categories_entities (id, alias, title, create_at, update_at) FROM stdin;
    public          postgres    false    223    B       ,          0    34135    locations_entities 
   TABLE DATA           �   COPY public.locations_entities (id, address1, address2, address3, city, zip_code, country, state, latitude, longitude, create_at, update_at) FROM stdin;
    public          postgres    false    220   C       )          0    34111    transactions_entities 
   TABLE DATA           V   COPY public.transactions_entities (id, transaction, create_at, update_at) FROM stdin;
    public          postgres    false    217   �C       ;           0    0    businesses_entities_id_seq    SEQUENCE SET     H   SELECT pg_catalog.setval('public.businesses_entities_id_seq', 8, true);
          public          postgres    false    214            <           0    0    categories_entities_id_seq    SEQUENCE SET     I   SELECT pg_catalog.setval('public.categories_entities_id_seq', 12, true);
          public          postgres    false    222            =           0    0    locations_entities_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.locations_entities_id_seq', 5, true);
          public          postgres    false    219            >           0    0    transactions_entities_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.transactions_entities_id_seq', 2, true);
          public          postgres    false    216            �           2606    34171 0   businesses_categories businesses_categories_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY public.businesses_categories
    ADD CONSTRAINT businesses_categories_pkey PRIMARY KEY (businesses_entity_id, categories_entity_id);
 Z   ALTER TABLE ONLY public.businesses_categories DROP CONSTRAINT businesses_categories_pkey;
       public            postgres    false    224    224            �           2606    34109 ,   businesses_entities businesses_entities_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.businesses_entities
    ADD CONSTRAINT businesses_entities_pkey PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.businesses_entities DROP CONSTRAINT businesses_entities_pkey;
       public            postgres    false    215            �           2606    34147 .   businesses_locations businesses_locations_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY public.businesses_locations
    ADD CONSTRAINT businesses_locations_pkey PRIMARY KEY (businesses_entity_id, locations_entity_id);
 X   ALTER TABLE ONLY public.businesses_locations DROP CONSTRAINT businesses_locations_pkey;
       public            postgres    false    221    221            �           2606    34123 4   businesses_transactions businesses_transactions_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY public.businesses_transactions
    ADD CONSTRAINT businesses_transactions_pkey PRIMARY KEY (businesses_entity_id, transactions_entity_id);
 ^   ALTER TABLE ONLY public.businesses_transactions DROP CONSTRAINT businesses_transactions_pkey;
       public            postgres    false    218    218            �           2606    34166 ,   categories_entities categories_entities_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.categories_entities
    ADD CONSTRAINT categories_entities_pkey PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.categories_entities DROP CONSTRAINT categories_entities_pkey;
       public            postgres    false    223            �           2606    34142 *   locations_entities locations_entities_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.locations_entities
    ADD CONSTRAINT locations_entities_pkey PRIMARY KEY (id);
 T   ALTER TABLE ONLY public.locations_entities DROP CONSTRAINT locations_entities_pkey;
       public            postgres    false    220            �           2606    34118 0   transactions_entities transactions_entities_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY public.transactions_entities
    ADD CONSTRAINT transactions_entities_pkey PRIMARY KEY (id);
 Z   ALTER TABLE ONLY public.transactions_entities DROP CONSTRAINT transactions_entities_pkey;
       public            postgres    false    217            �           2606    34172 @   businesses_categories fk_businesses_categories_businesses_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_categories
    ADD CONSTRAINT fk_businesses_categories_businesses_entity FOREIGN KEY (businesses_entity_id) REFERENCES public.businesses_entities(id);
 j   ALTER TABLE ONLY public.businesses_categories DROP CONSTRAINT fk_businesses_categories_businesses_entity;
       public          postgres    false    215    224    3205            �           2606    34177 @   businesses_categories fk_businesses_categories_categories_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_categories
    ADD CONSTRAINT fk_businesses_categories_categories_entity FOREIGN KEY (categories_entity_id) REFERENCES public.categories_entities(id);
 j   ALTER TABLE ONLY public.businesses_categories DROP CONSTRAINT fk_businesses_categories_categories_entity;
       public          postgres    false    224    3215    223            �           2606    34153 >   businesses_locations fk_businesses_locations_businesses_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_locations
    ADD CONSTRAINT fk_businesses_locations_businesses_entity FOREIGN KEY (businesses_entity_id) REFERENCES public.businesses_entities(id);
 h   ALTER TABLE ONLY public.businesses_locations DROP CONSTRAINT fk_businesses_locations_businesses_entity;
       public          postgres    false    215    3205    221            �           2606    34148 =   businesses_locations fk_businesses_locations_locations_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_locations
    ADD CONSTRAINT fk_businesses_locations_locations_entity FOREIGN KEY (locations_entity_id) REFERENCES public.locations_entities(id);
 g   ALTER TABLE ONLY public.businesses_locations DROP CONSTRAINT fk_businesses_locations_locations_entity;
       public          postgres    false    220    221    3211            �           2606    34124 D   businesses_transactions fk_businesses_transactions_businesses_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_transactions
    ADD CONSTRAINT fk_businesses_transactions_businesses_entity FOREIGN KEY (businesses_entity_id) REFERENCES public.businesses_entities(id);
 n   ALTER TABLE ONLY public.businesses_transactions DROP CONSTRAINT fk_businesses_transactions_businesses_entity;
       public          postgres    false    3205    215    218            �           2606    34129 F   businesses_transactions fk_businesses_transactions_transactions_entity    FK CONSTRAINT     �   ALTER TABLE ONLY public.businesses_transactions
    ADD CONSTRAINT fk_businesses_transactions_transactions_entity FOREIGN KEY (transactions_entity_id) REFERENCES public.transactions_entities(id);
 p   ALTER TABLE ONLY public.businesses_transactions DROP CONSTRAINT fk_businesses_transactions_transactions_entity;
       public          postgres    false    218    3207    217            0   "   x�3�4�2bc 6bS 6bs � �=... U9u      '   �  x��mk�0�?���O;F��hu0�o�֖��l$��l��F��_?�8ka�&�
D���x_y�� �k�G�&����fp�V��ݢ�`-(�i&Ls	<%I�/�U����|��n nh;�HtXX��'��������ް�5�榤fU&�c$o�dYV�~��o�_�����)O,gEI��ec8�:���\��1Dn�����Y���C�^t���"�o��Bإ�����N�Q�-N��y*Zp���c�C�`Q�!װ���f ]�Uѐ�����
�$ɰxI���B��?>+,H��pF�!I���\ ��������B^����M:�&7�>�̻d>��$��XGjE�9(
��8�1,��A�?-�`�R��򺾨m�q�@�r?j^}���O�|I���5���iHB���nzy���v7hݎ�I�u�:�vg�X��ws���?����&������\��B���; �����4P�ޔ�6�W	�� �0���l߶
I4�܏�j�P�3a�ǥ�`9i_��Fغ"]#P��5�J�G�8	�j�А"����t$!MW�0�lʓ�N�Ѡ�s��.��R0s9�+�rN�^���$J��]﹛�$�kɠ�OS:	����Ё�[p6�9������#���t���f@�<�����|'a+�t��5��X�T^ m��y      -   "   x�3�4�2bc 6bS 6bs � �=... T�m      *   "   x�3�4�2bc 6bS 6bs � �=... U9u      /   �   x���?k�0���O��KH��4i��I�JݱPd�L�c;�l���(1�2�p<=���$tJ��d��H)�!7y��2[���b��U�M}$F�oÉx��(/Ny�U����٠ͣ���YQ^}8r��hR<��q0Z�H�ݦ��g��m^�9/Pv�&��c��Qn>K�C�a�� �H-Ʈy�Ɏ�g� ������=�Oe��>����5�`T�%kܿ(f��$]�����Jn�s��(��,��      ,   �   x��ѱ
�@�9}��Ғ��rƭ���� ���`�V}}��
*.7$d���c`����=�: X�+n�� L$��7੊�L��*���q�xH:�2#?y���P�9.��k��,�Ȋ����ׁ��/�S|ud��8dp�\���)uXw�?u�QW��j��H3� ����~��b��:��!��VEQ� ��z      )   8   x�3�,�L�.-�4000�#3+#+m�$�gJjNfYjQ%���qqq g�     