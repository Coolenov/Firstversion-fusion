from aiogram import Bot, Dispatcher, executor, types
import markups
import naxApi
import buttons
from dotenv import load_dotenv
import os

load_dotenv()
API_TOKEN = os.getenv('API_TOKEN')

bot = Bot(token=API_TOKEN,parse_mode="HTML")
nax = naxApi.Nax()
dp = Dispatcher(bot)


@dp.message_handler(commands='start')
async def start_cmd_handler(message: types.Message):

	await message.answer("Pick source", reply_markup=markups.getStartMarkup())


@dp.callback_query_handler(buttons.source_callback.filter())
async def inline_kb_answer_callback_handler(query: types.CallbackQuery, callback_data:dict):
	content = nax.GetLastContentBySource(callback_data["src"])
	text = nax.GetTextForTelegramMessage(content)
	await query.message.edit_text(text,reply_markup=markups.getPaginationMarkup(content))


@dp.callback_query_handler(text='back_to_main_menu') 
async def inline_kb_answer_callback_handler(query: types.CallbackQuery):
	await start_cmd_handler(query.message)


@dp.callback_query_handler(buttons.next_post.filter())
async def next_callback_handler(query: types.CallbackQuery, callback_data: dict):
	content = nax.GetNextContent(callback_data["id"])
	text = nax.GetTextForTelegramMessage(content)
	await query.message.edit_text(text,reply_markup=markups.getPaginationMarkup(content))


@dp.callback_query_handler(buttons.prev_post.filter())
async def prev_callback_handler(query: types.CallbackQuery, callback_data: dict):
	content = nax.GetPreviousContent(callback_data["id"])
	text = nax.GetTextForTelegramMessage(content)
	await query.message.edit_text(text,reply_markup=markups.getPaginationMarkup(content))

if __name__ == '__main__':
	executor.start_polling(dp, skip_updates=True)