# Code generated by github.com/lolopinto/ent/ent, DO NOT edit. 

"""add column phone_number to table users
add unique constraint users_unique_phone_number

Revision ID: ec1f660b1524
Revises: fa1a57663f1c
Create Date: 2020-02-01 04:30:03.337402+00:00

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'ec1f660b1524'
down_revision = 'fa1a57663f1c'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('users', sa.Column('phone_number', sa.Text(), nullable=False))
    op.create_unique_constraint('users_unique_phone_number', 'users', ['phone_number'])
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_constraint('users_unique_phone_number', 'users', type_='unique')
    op.drop_column('users', 'phone_number')
    # ### end Alembic commands ###